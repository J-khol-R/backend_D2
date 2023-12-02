package repositories

import "github.com/J-khol-R/backend_D2/models"

type UserRepository interface {
	Create(user *models.User) error                       //crear una oferta
	Read(idUser int) (*models.User, error)                //obtener una oferta
	ReadByUsername(username string) (*models.User, error) //obtener una oferta
	Update(user *models.User) error                       //actualiza una oferta
}
