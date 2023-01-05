package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECT")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&UserModel{}, &TaskModel{})
	log.Print("Migrate!")
	return DB
}
