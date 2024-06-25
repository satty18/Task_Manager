package utils

import (
	"Task_Manager/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// AutoMigrate will create or update the tables to match the models
	err = DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
