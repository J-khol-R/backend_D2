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
	postulantRepository = repositories.NewPostulationGormRepository(db.Conn)
)

func GetOfferPostulations(w http.ResponseWriter, r *http.Request) {
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
	id := (variable["idOffer"])

	idOffer, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	postulations, err := postulantRepository.List(uint(idOffer))
	if err != nil {
		http.Error(w, "Error al traer las postulaciones: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(postulations)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreatePostulation(w http.ResponseWriter, r *http.Request) {
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

	var postulationReq models.Postulant
	err = json.NewDecoder(r.Body).Decode(&postulationReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = postulantRepository.Create(&postulationReq)
	if err != nil {
		http.Error(w, "Error al guardar oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("postulacion creada exitosamente"))
}

func DeletePostulation(w http.ResponseWriter, r *http.Request) {
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
	id := (variable["idPostulation"])

	idPostulation, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = postulantRepository.Delete(uint(idPostulation))
	if err != nil {
		http.Error(w, "Error al eliminar la postulacion: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("se elimino la postulacion correctamente"))

}
