package database

import (
	"api-alemao/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Phrases{})
}
