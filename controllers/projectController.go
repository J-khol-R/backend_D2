package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/models"
	"github.com/J-khol-R/backend_D2/repositories"
	"github.com/J-khol-R/backend_D2/security"
	"github.com/gorilla/mux"
)

var (
	projectRepository = repositories.NewProjectGormRepository(db.Conn)
	testRepository    = repositories.NewTestGormRepository(db.Conn)
)

func GetAllprojects(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	err := security.VerificateToken(token)
	if err != nil {
		http.Error(w, "Error al verificar token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err == security.ErrInvalidToken {
		http.Error(w, "token invalido: "+err.Error(), http.StatusForbidden)
		return
	}

	projects, err := projectRepository.List()
	if err != nil {
		http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var projectResponses []models.ProjectResponse
	for _, project := range projects {
		projectResponse := models.ProjectResponse{
			Id:          int(project.ID),
			Name:        project.Name,
			StartDate:   project.StartDate,
			FinishDate:  project.FinishDate,
			Description: project.Description,
		}

		tests, err := testRepository.List(project.ID)
		if err != nil {
			http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		projectResponse.Tests = tests
		projectResponses = append(projectResponses, projectResponse)
	}

	err = json.NewEncoder(w).Encode(projectResponses)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	err := security.VerificateToken(token)
	if err != nil {
		http.Error(w, "Error al verificar token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err == security.ErrInvalidToken {
		http.Error(w, "token invalido: "+err.Error(), http.StatusForbidden)
		return
	}

	var projectReq models.Project
	err = json.NewDecoder(r.Body).Decode(&projectReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = projectRepository.Create(&projectReq)
	if err != nil {
		http.Error(w, "Error al guardar proyecto: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("proyecto creado exitosamente"))
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	err := security.VerificateToken(token)
	if err != nil {
		http.Error(w, "Error al verificar token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err == security.ErrInvalidToken {
		http.Error(w, "token invalido: "+err.Error(), http.StatusForbidden)
		return
	}

	variable := mux.Vars(r)
	id := (variable["id"])

	isProject, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	project, err := projectRepository.Read(uint(isProject))
	if err != nil {
		http.Error(w, "Error al buscar el id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	projectResponse := models.ProjectResponse{
		Id:          int(project.ID),
		Name:        project.Name,
		StartDate:   project.StartDate,
		FinishDate:  project.FinishDate,
		Description: project.Description,
	}

	tests, err := testRepository.List(project.ID)
	if err != nil {
		http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	projectResponse.Tests = tests

	err = json.NewEncoder(w).Encode(projectResponse)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
