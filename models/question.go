package models

import "gorm.io/gorm"


type Question struct {
	gorm.Model
	Question string `json:"question" validate:"required"`
	UserID uint `json:"userId" validate:"required"`
}