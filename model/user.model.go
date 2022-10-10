package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           *string `json:"name" validate:"required,min=2,max=100"`
	HashedPassword *string `json:"-"`
	Email          *string `json:"email" validate:"email,required"`
	Products       []Product
}

type UserPrivate struct {
	User
	Password *string `json:"password" validate:"required,min=2,max=100"`
}
