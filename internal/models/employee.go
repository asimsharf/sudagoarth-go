package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
