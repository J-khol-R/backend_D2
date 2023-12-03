package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Titulo      string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	Precio      string `gorm:"not null"`
	Fecha_pub   string
	Fecha_exp   string
}

type CategoriaServicio struct {
	gorm.Model
	IDService int
	Service   Service `gorm:"foreignKey:IDService"`
	Categoria string
}

type ImagenServicio struct {
	gorm.Model
	IDService int
	Service   Service `gorm:"foreignKey:IDService"`
	ImagenURL string
}
