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

// // Create implementa el método Create de UserRepository
// func (r *OfferGormRepository) Create(user *User) error {
//     return r.db.Create(user).Error
// }

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

// // Update implementa el método Update de UserRepository
// func (r *OfferGormRepository) Update(user *User) error {
//     return r.db.Save(user).Error
// }

// // Delete implementa el método Delete de UserRepository
// func (r *OfferGormRepository) Delete(id uint) error {
//     return r.db.Delete(&User{}, id).Error
// }

// lista todas las ofertas acorde a una condicion
func (r *OfferGormRepository) List(condition map[string]interface{}) ([]models.Offer, error) {
	var offers []models.Offer
	if err := r.db.Where(condition).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}
