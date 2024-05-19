package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name      string
	IsPrivate bool
	PassCode  string
	PlayerIDs []uint `gorm:"-"` // Not stored in DB, managed in Redis
}
