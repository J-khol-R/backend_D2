package db

import "github.com/J-khol-R/backend_D2/models"

func MigrateModels() {
	Conn.AutoMigrate(&models.User{},
		&models.CategoriaOferta{},
		&models.ImagenOferta{},
		&models.Score{},
		&models.Postulant{},
		&models.Offer{})
}
