package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     *string `json:"name" validate:"required,min=2,max=100"`
	Password *string `json:"password" validate:"required,min=2,max=100"`
	Email    *string `json:"email" validate:"email,required"`
}
