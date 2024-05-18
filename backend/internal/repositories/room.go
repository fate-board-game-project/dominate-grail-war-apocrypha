package repositories

import (
	"backend/internal/models"
	"backend/pkg/db"
)

func CreateRoom(room *models.Room) error {
	return db.DB.Create(room).Error
}

func GetRoomByID(id uint) (*models.Room, error) {
	var room models.Room
	err := db.DB.First(&room, id).Error
	return &room, err
}
