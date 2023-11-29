package models

import (
	"gorm.io/gorm"
)

type Postulant struct {
	gorm.Model
	IDOffer string `gorm:"foreignKey:IDOffer"`
	IDUser  string `gorm:"foreignKey:IDUser"`
	Estado  string
}
