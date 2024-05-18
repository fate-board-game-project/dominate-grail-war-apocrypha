package repositories

import (
	"backend/internal/models"
	"backend/pkg/db"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
