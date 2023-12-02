package models

import (
	"gorm.io/gorm"
)

type Offer struct {
	gorm.Model
	IDServicio     int    `gorm:"foreignKey:IDService"`
	IDOfertante    int    `gorm:"foreignKey:IDUser"`
	IDCalificacion int    `gorm:"unique;foreignKey:IDScore"`
	Estado         string `gorm:"not null"`
}
