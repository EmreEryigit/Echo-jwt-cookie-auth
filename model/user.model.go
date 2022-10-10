package model

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

func (u *UserPrivate) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(*u.Password), 8)
	if err != nil {
		log.Panic(err)
	}
	str := string(hash)
	u.HashedPassword = &str
}
