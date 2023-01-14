package db

import (
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ChatID   uint64
	Username string
}

type TaskModel struct {
	gorm.Model
	ChatID        uint64
	Name          string
	Priority      string
	Status        string
	Notification  time.Time
	TakenToWorkAt time.Time `gorm:"default:null"`
	CompletedAt   time.Time `gorm:"default:null"`
}
