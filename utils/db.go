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
	{ /*
				dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_NAME"),
			)

			DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatal("Failed to connect to database:", err)
			}
		*/
	}

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Migrate models
	err = DB.AutoMigrate(&models.User{}, &models.Role{})
	if err != nil {
		log.Fatal("Failed to migrate models!", err)
	}
}
