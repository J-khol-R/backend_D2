package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nombre     string `gorm:"not null"`
	Telefono   string
	Ciudad     string
	Correo     string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	Contraseña string `gorm:"not null"`
}
