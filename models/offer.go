package models

import (
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	Titulo      string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	Precio      string
	Fecha_exp   string
	Categorias  []CategoriaOferta `gorm:"foreignKey:IDOffer"`
	ImagenesURL []ImagenOferta    `gorm:"foreignKey:IDOffer"`
	IDOfertante uint              `gorm:"not null"`
	Ofertante   User              `gorm:"foreignKey:IDOfertante"`
	Estado      string            `gorm:"not null"`
}

type OfferRequest struct {
	Id          int      `json:"id"`
	Titulo      string   `json:"titulo"`
	Descripcion string   `json:"descripcion"`
	Precio      string   `json:"precio"`
	Fecha_exp   string   `json:"fecha_exp"`
	Categorias  []string `json:"categorias"`
	ImagenesURL []string `json:"imagenes_url"`
	IDOfertante uint     `json:"id_ofertante"`
	Estado      string   `json:"estado"`
}

type CategoriaOferta struct {
	gorm.Model
	IDOffer   int
	Offer     Offer `gorm:"foreignKey:IDOffer"`
	Categoria string
}

type ImagenOferta struct {
	gorm.Model
	IDOffer   int
	Offer     Offer `gorm:"foreignKey:IDOffer"`
	ImagenURL string
}
