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
	userRepository = repositories.NewUserGormRepository(db.Conn)
)

func GetSesion(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Error al decodificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := userRepository.ReadByUsername(login.Username)
	if err != nil {
		http.Error(w, "Error al buscar el usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = security.CheckPassword(login.Password, user.Contraseña)
	if err != nil {
		http.Error(w, "contraseña incorrecta: "+err.Error(), http.StatusForbidden)
		return
	}

	token, err := security.GenerateToken(int(user.ID))
	if err != nil {
		http.Error(w, "Error creando token de acceso: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tokenResponse := models.TokenResponse{
		Token: token,
	}

	err = json.NewEncoder(w).Encode(tokenResponse)
	if err != nil {
		http.Error(w, "Error al codificar la respuesta: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
