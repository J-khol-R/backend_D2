package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Nro_estrellas     string `gorm:"not null"`
	Comentarios       string
	Fecha             string
	Calificador       string `gorm:"foreignKey:IDUser"`
	Usuario_calficado string `gorm:"foreignKey:IDUser"`
}
