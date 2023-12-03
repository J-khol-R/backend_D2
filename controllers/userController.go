package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/models"
	"github.com/J-khol-R/backend_D2/repositories"
	"github.com/J-khol-R/backend_D2/security"
)

var (
	UserRepository = repositories.NewUserGormRepository(db.Conn)
)

func GetUser(w http.ResponseWriter, r *http.Request) {
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	hashedPassword, err := security.HashPassword(user.Contraseña)
	if err != nil {
		return
	}

	user.Contraseña = hashedPassword

	err = userRepository.Create(&user)
	if err != nil {
		http.Error(w, "Error alguardar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("usuario creado exitosamente"))
}
