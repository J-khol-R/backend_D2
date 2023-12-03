package repositories

import (
	"github.com/J-khol-R/backend_D2/models"
	"gorm.io/gorm"
)

type OfferGormRepository struct {
	db *gorm.DB
}

func NewOfferGormRepository(db *gorm.DB) *OfferGormRepository {
	return &OfferGormRepository{db}
}

// Read implementa el método Read de UserRepository
func (r *OfferGormRepository) Read(id uint) (*models.Offer, error) {
	var offer models.Offer

	result := r.db.Preload("Servicio").
		Preload("Ofertante").
		First(&offer, id)

	if result.Error != nil {
		return &offer, result.Error
	}
	return &offer, nil
}

// Delete implementa el método Delete de UserRepository
func (r *OfferGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Offer{}, id).Error
}

// lista todas las ofertas acorde a una condicion
func (r *OfferGormRepository) List(condition map[string]interface{}) ([]models.Offer, error) {
	var offers []models.Offer
	if err := r.db.Where(condition).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}

func (r *OfferGormRepository) ListCreatedUserOffers(idUser uint) ([]models.Offer, error) {
	var offers []models.Offer
	if err := r.db.Where("id_ofertante = ?", idUser).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}

func (r *OfferGormRepository) ListPostulatedUserOffers(idUser uint) ([]models.Offer, error) {
	var offersPostulant []models.Offer
	var postulant []models.Postulant

	if err := r.db.Where("IDUser = ?", idUser).Find(&postulant).Error; err != nil {
		return nil, err
	}

	for _, postulacion := range postulant {
		oferta := models.Offer{}
		if err := r.db.First(&oferta, postulacion.IDOffer).Error; err != nil {
			return nil, err
		}
		offersPostulant = append(offersPostulant, oferta)
	}
	return offersPostulant, nil
}
