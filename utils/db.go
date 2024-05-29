package utils

import (
	"log"

	"github.com/aleksanderpalamar/rbac-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Migrate models
	DB.AutoMigrate(&models.User{}, &models.Role{})
}
