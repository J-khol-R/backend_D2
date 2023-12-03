package models

import (
	"gorm.io/gorm"
)

type Postulant struct {
	gorm.Model
	IDOffer uint
	Offer   Offer `gorm:"foreignKey:IDOffer"`
	IDUser  uint
	User    User `gorm:"foreignKey:IDUser"`
	Estado  string
}
