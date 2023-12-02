package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nombre     string `gorm:"not null"`
	Telefono   string `gorm:"not null"`
	Ciudad     string `gorm:"not null"`
	Correo     string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	Contraseña string `gorm:"not null"`
}
