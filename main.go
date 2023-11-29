package main

import (
	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/models"
)

func main() {
	db.Conn.AutoMigrate(&models.User{},
		&models.Service{},
		&models.CategoriaServicio{},
		&models.ImagenServicio{},
		&models.Score{},
		&models.Postulant{},
		&models.Offer{})

	nuevoServicio := models.Service{
		Titulo:      "hola",
		Descripcion: "muy buen trabajo",
		Precio:      "1",
		Fecha_pub:   "3",
		Fecha_exp:   "4",
	}

	db.Conn.Create(&nuevoServicio)
}
