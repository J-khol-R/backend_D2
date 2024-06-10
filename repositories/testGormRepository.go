package repositories

import (
	"github.com/J-khol-R/backend_D2/models"
	"gorm.io/gorm"
)

type TestGormRepository struct {
	db *gorm.DB
}

func NewTestGormRepository(db *gorm.DB) *TestGormRepository {
	return &TestGormRepository{db}
}

func (r *TestGormRepository) Create(Test *models.Test) error {
	return r.db.Create(Test).Error
}

func (r *TestGormRepository) List(idProject uint) ([]models.Test, error) {
	var tests []models.Test

	if err := r.db.Where("id_project = ?", idProject).Find(&tests).Error; err != nil {
		return nil, err
	}

	return tests, nil
}

func (r *TestGormRepository) Read(id uint) (*models.Test, error) {
	var test models.Test
	if err := r.db.First(&test, id).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

func (r *TestGormRepository) Update(idTest uint, newTest models.Test) error {
	var test models.Test
	if err := r.db.First(&test, idTest).Error; err != nil {
		return err
	}

	if newTest.Title != "" {
		test.Title = newTest.Title
	}
	if newTest.PublicationDate != "" {
		test.PublicationDate = newTest.PublicationDate
	}
	if newTest.LimitDate != "" {
		test.LimitDate = newTest.LimitDate
	}
	if newTest.Priority != "" {
		test.Priority = newTest.Priority
	}
	if newTest.Fase != "" {
		test.Fase = newTest.Fase
	}
	if newTest.Evidence != "" {
		test.Evidence = newTest.Evidence
	}
	if newTest.Description != "" {
		test.Description = newTest.Description
	}
	if newTest.Expectations != "" {
		test.Expectations = newTest.Expectations
	}
	if newTest.Steps != "" {
		test.Steps = newTest.Steps
	}
	if newTest.Report != "" {
		test.Report = newTest.Report
	}
	if newTest.Responsable != "" {
		test.Responsable = newTest.Responsable
	}
	if newTest.State != "" {
		test.State = newTest.State
	}

	if err := r.db.Save(&test).Error; err != nil {
		return err
	}

	return nil
}
