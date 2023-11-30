package db

import "github.com/J-khol-R/backend_D2/models"

func MigrateModels() {
	Conn.AutoMigrate(&models.User{},
		&models.Service{},
		&models.CategoriaServicio{},
		&models.ImagenServicio{},
		&models.Score{},
		&models.Postulant{},
		&models.Offer{})
}
