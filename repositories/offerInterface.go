package repositories

import "github.com/J-khol-R/backend_D2/models"

type OfferRepository interface {
	Create(offer *models.Offer) error                              //crear una oferta
	Read(idOffer int) (*models.Offer, error)                       //obtener una oferta
	Update(offer *models.Offer) error                              //actualiza una oferta
	Delete(idOffer int) error                                      //eliminar una oferta
	List(condition map[string]interface{}) ([]models.Offer, error) //listar todas las ofertas disponibles
	ListCreatedUserOffers(idUser int) ([]models.Offer, error)      //listar las ofertas creadas por un usuario
	ListPostulatedUserOffers(idUser int) ([]models.Offer, error)   //listar las ofertas a las que se ha postulado un usuario
}
