package database

import (
	"echojwt/model"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"))
	if err != nil {
		panic("failed to connect to DB")
	}

	db.AutoMigrate(&model.User{})
	fmt.Println("Connected to db")

	return db
}
