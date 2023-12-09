package models

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	IdUser int    `json:"id_user"`
	Token  string `json:"token"`
}
