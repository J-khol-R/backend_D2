package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

type EnvConfig struct {
	Host        string
	Port        string
	DbName      string
	RolName     string
	RolPassword string
}

func GetEnvConfig() EnvConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envConfig := EnvConfig{
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DbName:      os.Getenv("DB_NAME"),
		RolName:     os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
	}

	return envConfig
}

func OpenDbConnection() *gorm.DB {
	envConfig := GetEnvConfig()

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
