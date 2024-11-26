package db

import (
	"backend/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupPostgres(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = DB.AutoMigrate(&models.User{}, &models.Room{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
