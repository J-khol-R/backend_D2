package security

import (
	"errors"
	"time"

	"github.com/J-khol-R/backend_D2/services"
	"github.com/golang-jwt/jwt/v4"
)

var (
	envConfig       = services.GetEnvConfig()
	ErrInvalidToken = errors.New("token invalido")
)

func GenerateToken(idUser int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": idUser,
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // Token expira en 2 horas
	})

	tokenString, err := token.SignedString([]byte(envConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerificateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(envConfig.SecretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	return nil
}
