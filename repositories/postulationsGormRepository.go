package repositories

import (
	"errors"

	"github.com/J-khol-R/backend_D2/models"
	"gorm.io/gorm"
)

type PostulationGormRepository struct {
	db *gorm.DB
}

func NewPostulationGormRepository(db *gorm.DB) *PostulationGormRepository {
	return &PostulationGormRepository{db}
}

func (r *PostulationGormRepository) Create(postulation *models.Postulant) error {
	var existingPostulation models.Postulant
	if err := r.db.Where("id_user = ? AND id_offer = ?", postulation.IDUser, postulation.IDOffer).First(&existingPostulation).Error; err == nil {
		return errors.New("el usuario ya está postulado a esta oferta")
	}

	return r.db.Create(postulation).Error
}

func (r *PostulationGormRepository) Delete(id uint) error {
	return r.db.Delete(&models.Postulant{}, id).Error
}

func (r *PostulationGormRepository) List(idOffer uint) ([]models.Postulant, error) {
	var postulations []models.Postulant
	if err := r.db.Preload("Offer").Preload("User").Where("id_offer = ?", idOffer).Find(&postulations).Error; err != nil {
		return nil, err
	}
	return postulations, nil
}
