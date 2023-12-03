package models

import (
	"gorm.io/gorm"
)

// type Offer struct {
// 	gorm.Model
// 	IDServicio     int    `gorm:"foreignKey:IDService"`
// 	IDOfertante    int    `gorm:"foreignKey:IDUser"`
// 	IDCalificacion int    `gorm:"unique;foreignKey:IDScore"`
// 	Estado         string `gorm:"not null"`
// }

type Offer struct {
	gorm.Model
	IDServicio  uint
	Servicio    Service `gorm:"foreignKey:IDServicio"`
	IDOfertante uint
	Ofertante   User   `gorm:"foreignKey:IDOfertante"`
	Estado      string `gorm:"not null"`
}
