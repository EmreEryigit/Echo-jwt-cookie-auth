package database

import (
	"echojwt/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sqlite.db"))
	if err != nil {
		panic("failed to connect to DB")
	}

	db.AutoMigrate(&model.User{})
	return db
}
