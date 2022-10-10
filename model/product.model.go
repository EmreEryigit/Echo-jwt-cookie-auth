package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"required,min=2,max=100"`
	Price       int    `json:"price" validate:"required"`
	UserID      uint
}
