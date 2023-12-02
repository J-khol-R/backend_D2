package repositories

import "github.com/J-khol-R/backend_D2/models"

type PostulationRepository interface {
	Create(postulation *models.Postulant) error                  //crear una postulacion
	Delete(idPOstulation, idUser int) error                      //eliminar una postulacion
	ListUserPostulations(idUser int) ([]models.Postulant, error) //listar todas las postulaciones de un usuario
}
