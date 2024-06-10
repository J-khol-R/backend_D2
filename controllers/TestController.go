package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/models"
	"github.com/J-khol-R/backend_D2/repositories"
	"github.com/J-khol-R/backend_D2/security"
	"github.com/gorilla/mux"
)

var (
	testsRepository = repositories.NewTestGormRepository(db.Conn)
)

func GetTestById(w http.ResponseWriter, r *http.Request) {
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

	idTest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	test, err := testsRepository.Read(uint(idTest))
	if err != nil {
		http.Error(w, "Error al traer las postulaciones: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(test)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateTest(w http.ResponseWriter, r *http.Request) {
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

	var testReq models.Test
	err = json.NewDecoder(r.Body).Decode(&testReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fecha := time.Now()
	testReq.PublicationDate = fecha.Format("2006-01-02")

	err = testsRepository.Create(&testReq)
	if err != nil {
		http.Error(w, "Error al guardar test: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("test creado exitosamente"))
}

func UpdateTest(w http.ResponseWriter, r *http.Request) {
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

	idTest, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	var testReq models.Test
	err = json.NewDecoder(r.Body).Decode(&testReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = testRepository.Update(uint(idTest), testReq)
	if err != nil {
		http.Error(w, "Error al actualizar test: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("se actualizo el test correctamente"))
}

// func DeletePostulation(w http.ResponseWriter, r *http.Request) {
// 	token := r.Header.Get("Authorization")

// 	err := security.VerificateToken(token)
// 	if err != nil {
// 		http.Error(w, "Error al verificar token: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if err == security.ErrInvalidToken {
// 		http.Error(w, "token invalido: "+err.Error(), http.StatusForbidden)
// 		return
// 	}

// 	variable := mux.Vars(r)
// 	id := (variable["idPostulation"])

// 	idPostulation, err := strconv.Atoi(id)
// 	if err != nil {
// 		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = testsRepository.Delete(uint(idPostulation))
// 	if err != nil {
// 		http.Error(w, "Error al eliminar la postulacion: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("se elimino la postulacion correctamente"))

// }
