package repositories

import "github.com/J-khol-R/backend_D2/models"

type ScoreRepository interface {
	Create(calification *models.Score) error                  //crear una calificacion
	ListUserCalifications(idUser int) ([]models.Score, error) //listar todas las calificaciones de un usuario
}
