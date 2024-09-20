package database

import (
	"golang-backend-test/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database/bayarind.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.User{})
	return db, nil
}
