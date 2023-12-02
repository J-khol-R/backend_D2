package db

import (
	"fmt"

	"github.com/J-khol-R/backend_D2/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func OpenDbConnection() *gorm.DB {
	envConfig := services.GetEnvConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envConfig.Host, envConfig.Port, envConfig.RolName, envConfig.RolPassword, envConfig.DbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos")
	}

	return db
}

func init() {
	Conn = OpenDbConnection()
}
