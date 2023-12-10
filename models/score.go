package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Nro_estrellas      int `gorm:"not null"`
	Comentarios        string
	Fecha              string
	Calificador        int
	UserCalificador    User `gorm:"foreignKey:Calificador"`
	Usuario_calificado int
	UserCalificado     User `gorm:"foreignKey:Usuario_calificado"`
	IDOferta           int
	Offer              Offer `gorm:"foreignKey:IDOferta" json:"-"`
}
