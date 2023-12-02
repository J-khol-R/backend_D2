package models

type Login struct {
	Username string
	Password string
}

type TokenResponse struct {
	Token string
}
