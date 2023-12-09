package repositories

import (
	"github.com/J-khol-R/backend_D2/db"
	"github.com/J-khol-R/backend_D2/models"
	"github.com/J-khol-R/backend_D2/security"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db}
}

func (r *UserGormRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserGormRepository) Read(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) ReadByUsername(username string) (*models.User, error) {
	var user models.User
	if err := db.Conn.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserGormRepository) Update(idUser int, newUser models.User) error {
	var user models.User
	if err := r.db.First(&user, idUser).Error; err != nil {
		return err
	}

	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Phone != "" {
		user.Phone = newUser.Phone
	}
	if newUser.City != "" {
		user.City = newUser.City
	}
	if newUser.Password != "" {
		password, err := security.HashPassword(newUser.Password)
		if err != nil {
			return err
		}
		user.Password = password
	}
	if newUser.ImageURL != "" {
		user.ImageURL = newUser.ImageURL
	}

	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
