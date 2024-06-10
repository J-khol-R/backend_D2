package db

import "github.com/J-khol-R/backend_D2/models"

func MigrateModels() {
	Conn.AutoMigrate(&models.User{},
		&models.User{},
		&models.Project{},
		&models.Test{})
}
