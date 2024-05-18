package services

import (
	"backend/internal/models"
	"backend/pkg/db"
	"backend/utils"
	"errors"
)

func RegisterUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := models.User{Username: username, Password: hashedPassword}
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(username, password string) (string, error) {
	var user models.User
	if err := db.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
