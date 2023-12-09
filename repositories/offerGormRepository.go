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

	result := r.db.Preload("Ofertante").Preload("Categorias").Preload("ImagenesURL").First(&offer, id)

	if result.Error != nil {
		return &offer, result.Error
	}
	return &offer, nil
}

func (r *OfferGormRepository) Create(offer *models.Offer) error {
	return r.db.Create(offer).Error
}

func (r *OfferGormRepository) CreateCategory(offerCategory *models.CategoriaOferta) error {
	return r.db.Create(offerCategory).Error
}

func (r *OfferGormRepository) CreateImage(offerImage *models.ImagenOferta) error {
	return r.db.Create(offerImage).Error
}

// Delete implementa el método Delete de UserRepository
func (r *OfferGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Offer{}, id).Error
}

func (r *OfferGormRepository) Update(idOffer int, newOffer models.OfferRequest) error {
	var offer models.Offer
	if err := r.db.First(&offer, idOffer).Error; err != nil {
		return err
	}

	if newOffer.Titulo != "" {
		offer.Titulo = newOffer.Titulo
	}
	if newOffer.Descripcion != "" {
		offer.Descripcion = newOffer.Descripcion
	}
	if newOffer.Precio != "" {
		offer.Precio = newOffer.Precio
	}
	if newOffer.Estado != "" {
		offer.Estado = newOffer.Estado
	}

	if err := r.db.Save(&offer).Error; err != nil {
		return err
	}

	return nil
}

// lista todas las ofertas acorde a una condicion
func (r *OfferGormRepository) List(condition map[string]interface{}) ([]models.Offer, error) {
	var offers []models.Offer
	if err := r.db.Preload("Ofertante").Preload("Categorias").Preload("ImagenesURL").Where(condition).Find(&offers).Error; err != nil {
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

	if err := r.db.Where("id_user = ?", idUser).Find(&postulant).Error; err != nil {
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
