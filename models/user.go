package models

import "gorm.io/gorm"


type User struct{
	gorm.Model
	Name string `gorm:"column:name" json:"name" validate:"required"`
	Email string `gorm:"column:email" json:"email" validate:"required,email"`
	Questions []Question `json:"questions"`
}