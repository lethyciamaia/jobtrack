package database

import (
	"fmt"
	"jobtrack-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.Application{})
	if err != nil {
		panic("Failed to migrate database")
	}

	DB = db
	fmt.Println("Successfully connected to the database!")
}
