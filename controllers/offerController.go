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

	var offersreq []models.OfferRequest
	for _, offer := range offers {
		offerReq := models.OfferRequest{
			Id:          int(offer.ID),
			Titulo:      offer.Titulo,
			Descripcion: offer.Descripcion,
			Precio:      offer.Precio,
			Fecha_exp:   offer.Fecha_exp,
			IDOfertante: offer.IDOfertante,
			Estado:      offer.Estado,
		}

		for _, categoria := range offer.Categorias {
			offerReq.Categorias = append(offerReq.Categorias, categoria.Categoria)
		}

		for _, imagen := range offer.ImagenesURL {
			offerReq.ImagenesURL = append(offerReq.ImagenesURL, imagen.ImagenURL)
		}

		offersreq = append(offersreq, offerReq)
	}

	err = json.NewEncoder(w).Encode(offersreq)
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

	offerReq := models.OfferRequest{
		Id:          int(offer.ID),
		Titulo:      offer.Titulo,
		Descripcion: offer.Descripcion,
		Precio:      offer.Precio,
		Fecha_exp:   offer.Fecha_exp,
		IDOfertante: offer.IDOfertante,
		Estado:      offer.Estado,
	}

	for _, categoria := range offer.Categorias {
		offerReq.Categorias = append(offerReq.Categorias, categoria.Categoria)
	}

	for _, imagen := range offer.ImagenesURL {
		offerReq.ImagenesURL = append(offerReq.ImagenesURL, imagen.ImagenURL)
	}

	err = json.NewEncoder(w).Encode(offerReq)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateOffer(w http.ResponseWriter, r *http.Request) {
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

	var offerReq models.OfferRequest
	err = json.NewDecoder(r.Body).Decode(&offerReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	offer := models.Offer{
		Titulo:      offerReq.Titulo,
		Descripcion: offerReq.Descripcion,
		Precio:      offerReq.Precio,
		Fecha_exp:   offerReq.Fecha_exp,
		IDOfertante: offerReq.IDOfertante,
		Estado:      offerReq.Estado,
	}

	err = offerRepository.Create(&offer)
	if err != nil {
		http.Error(w, "Error al guardar oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, value := range offerReq.Categorias {
		category := models.CategoriaOferta{
			IDOffer:   int(offer.ID),
			Categoria: value,
		}

		err = offerRepository.CreateCategory(&category)
		if err != nil {
			http.Error(w, "Error al guardar categoria: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	for _, value := range offerReq.ImagenesURL {
		image := models.ImagenOferta{
			IDOffer:   int(offer.ID),
			ImagenURL: value,
		}

		err = offerRepository.CreateImage(&image)
		if err != nil {
			http.Error(w, "Error al guardar imagen: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("oferta creado exitosamente"))
}

func UpdateOffer(w http.ResponseWriter, r *http.Request) {
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

	var offerReq models.OfferRequest
	err = json.NewDecoder(r.Body).Decode(&offerReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = offerRepository.Update(idOffer, offerReq)
	if err != nil {
		http.Error(w, "Error al actualizar oferta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("se actualizo la oferta correctamente"))
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

	var offersreq []models.OfferRequest
	for _, offer := range offers {
		offerReq := models.OfferRequest{
			Id:          int(offer.ID),
			Titulo:      offer.Titulo,
			Descripcion: offer.Descripcion,
			Precio:      offer.Precio,
			Fecha_exp:   offer.Fecha_exp,
			IDOfertante: offer.IDOfertante,
			Estado:      offer.Estado,
		}

		for _, categoria := range offer.Categorias {
			offerReq.Categorias = append(offerReq.Categorias, categoria.Categoria)
		}

		for _, imagen := range offer.ImagenesURL {
			offerReq.ImagenesURL = append(offerReq.ImagenesURL, imagen.ImagenURL)
		}

		offersreq = append(offersreq, offerReq)
	}

	err = json.NewEncoder(w).Encode(offersreq)
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

	var offersreq []models.OfferRequest
	for _, offer := range offers {
		offerReq := models.OfferRequest{
			Id:          int(offer.ID),
			Titulo:      offer.Titulo,
			Descripcion: offer.Descripcion,
			Precio:      offer.Precio,
			Fecha_exp:   offer.Fecha_exp,
			IDOfertante: offer.IDOfertante,
			Estado:      offer.Estado,
		}

		for _, categoria := range offer.Categorias {
			offerReq.Categorias = append(offerReq.Categorias, categoria.Categoria)
		}

		for _, imagen := range offer.ImagenesURL {
			offerReq.ImagenesURL = append(offerReq.ImagenesURL, imagen.ImagenURL)
		}

		offersreq = append(offersreq, offerReq)
	}

	err = json.NewEncoder(w).Encode(offersreq)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
