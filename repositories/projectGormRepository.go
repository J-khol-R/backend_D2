package repositories

import (
	"github.com/J-khol-R/backend_D2/models"
	"gorm.io/gorm"
)

type ProjectGormRepository struct {
	db *gorm.DB
}

func NewProjectGormRepository(db *gorm.DB) *ProjectGormRepository {
	return &ProjectGormRepository{db}
}

func (r *ProjectGormRepository) List() ([]models.Project, error) {
	var projects []models.Project

	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectGormRepository) Create(projects *models.Project) error {
	return r.db.Create(projects).Error
}
