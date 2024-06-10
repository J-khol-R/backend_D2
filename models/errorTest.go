package models

import (
	"gorm.io/gorm"
)

type ErrorTest struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Responsable     string `gorm:"not null"`
	Status          string `gorm:"not null"`
	UnExecutedCases string
}
