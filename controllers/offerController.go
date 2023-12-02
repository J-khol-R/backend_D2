package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/repositories"
	"github.com/J-khol-R/backend_D2/security"
)

var (
	offerRepository = repositories.NewOfferGormRepository(db.Conn)
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
func CreateOffer(w http.ResponseWriter, r *http.Request) {
}
func GetOffer(w http.ResponseWriter, r *http.Request) {
}
func UpdateOffer(w http.ResponseWriter, r *http.Request) {
}
func DeleteOffer(w http.ResponseWriter, r *http.Request) {
}
func GetCreatedUserOffers(w http.ResponseWriter, r *http.Request) {
}
func GetPostutatedUserOffers(w http.ResponseWriter, r *http.Request) {
}
