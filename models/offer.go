package models

import (
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	IDServicio     string `gorm:"foreignKey:IDService"`
	IDOfertante    string `gorm:"foreignKey:IDUser"`
	IDCalificacion string `gorm:"foreignKey:IDScore"`
	Estado         string `gorm:"not null"`
}
