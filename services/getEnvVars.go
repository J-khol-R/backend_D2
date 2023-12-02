package services

import (
	"os"

	"github.com/J-khol-R/backend_D2/models"
	"github.com/joho/godotenv"
)

func GetEnvConfig() models.EnvConfig {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envConfig := models.EnvConfig{
		Host:        os.Getenv("HOST"),
		Port:        os.Getenv("PORT"),
		DbName:      os.Getenv("DB_NAME"),
		RolName:     os.Getenv("ROL_NAME"),
		RolPassword: os.Getenv("ROL_PASSWORD"),
		SecretKey:   os.Getenv("SECRET_KEY"),
	}

	return envConfig
}
