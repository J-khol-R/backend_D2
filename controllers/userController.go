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
	UserRepository = repositories.NewUserGormRepository(db.Conn)
)

func GetUser(w http.ResponseWriter, r *http.Request) {
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

	idUser, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userRepository.Read(uint(idUser))
	if err != nil {
		http.Error(w, "Error al buscar el id: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	idUser, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error al convertir el id: "+err.Error(), http.StatusBadRequest)
		return
	}

	var userReq models.User
	err = json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = UserRepository.Update(idUser, userReq)
	if err != nil {
		http.Error(w, "Error al actualizar usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("usuario actualizado exitosamente"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return
	}

	user.Password = hashedPassword

	err = userRepository.Create(&user)
	if err != nil {
		http.Error(w, "Error alguardar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("usuario creado exitosamente"))
}
