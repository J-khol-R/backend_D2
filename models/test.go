package models

import (
	"gorm.io/gorm"
)

type Test struct {
	gorm.Model
	IDProject       string
	Project         Project `gorm:"foreignKey:IDProject" json:"-"`
	Title           string  `gorm:"not null"`
	PublicationDate string  `gorm:"not null"`
	LimitDate       string  `gorm:"not null"`
	Priority        string  `gorm:"not null"`
	Fase            string  `gorm:"not null"`
	Evidence        string
	Description     string `gorm:"not null"`
	Expectations    string `gorm:"not null"`
	Steps           string `gorm:"not null"`
	Report          string
	Responsable     string
	State           string `gorm:"not null"`
}
