package repositories

import (
	"github.com/J-khol-R/backend_D2/models"
	"gorm.io/gorm"
)

type ScoreGormRepository struct {
	db *gorm.DB
}

func NewScoreGormRepository(db *gorm.DB) *ScoreGormRepository {
	return &ScoreGormRepository{db}
}

func (r *ScoreGormRepository) List(idUser uint) ([]models.Score, error) {
	var calificaciones []models.Score

	if err := r.db.Where("usuario_calificado = ?", idUser).Find(&calificaciones).Error; err != nil {
		return nil, err
	}

	return calificaciones, nil
}

func (r *ScoreGormRepository) Create(calificaciones *models.Score) error {
	return r.db.Create(calificaciones).Error
}
