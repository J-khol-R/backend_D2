package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/repositories"
	"github.com/J-khol-R/backend_D2/security"
	"github.com/gorilla/mux"
)

var (
	offerRepository   = repositories.NewOfferGormRepository(db.Conn)
	serviceRepository = repositories.NewServiceGormRepository(db.Conn)
)

func GetAllOffers(w http.ResponseWriter, r *http.Request) {
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

	condition := map[string]interface{}{"estado": "Disponible"}
	offers, err := offerRepository.List(condition)
	if err != nil {
		http.Error(w, "Error al obtener datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(offers)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetOffer(w http.ResponseWriter, r *http.Request) {
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

	idOffer, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	offer, err := offerRepository.Read(uint(idOffer))
	if err != nil {
		http.Error(w, "Error al traer la oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(offer)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// se elimina la oferta y el servicio asociado
func DeleteOffer(w http.ResponseWriter, r *http.Request) {
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

	idOffer, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	offer, err := offerRepository.Read(uint(idOffer))
	if err != nil {
		http.Error(w, "Error al leer la oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = offerRepository.Delete(offer.ID)
	if err != nil {
		http.Error(w, "Error al eliminar la oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = serviceRepository.Delete(offer.IDServicio)
	if err != nil {
		http.Error(w, "Error al eliminar el servicio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("se elimino la oferta correctamente"))

}

func GetCreatedUserOffers(w http.ResponseWriter, r *http.Request) {
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
	id := (variable["idUser"])

	idUser, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	offers, err := offerRepository.ListCreatedUserOffers(uint(idUser))
	if err != nil {
		http.Error(w, "Error al obtener ofertas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(offers)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
func GetPostutatedUserOffers(w http.ResponseWriter, r *http.Request) {
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
	id := (variable["idUser"])

	idUser, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	offers, err := offerRepository.ListPostulatedUserOffers(uint(idUser))
	if err != nil {
		http.Error(w, "Error al obtener ofertas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(offers)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
