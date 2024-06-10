package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name        string `gorm:"not null"`
	StartDate   string `gorm:"not null"`
	FinishDate  string `gorm:"not null"`
	Description string `gorm:"not null"`
}

type ProjectResponse struct {
	Id          int
	Name        string
	StartDate   string
	FinishDate  string
	Description string
	Tests       []Test
}
