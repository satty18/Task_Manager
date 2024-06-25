package utils

import (
	"Task_Manager/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestInitDB(t *testing.T) {
	var err error
	DB, err = gorm.Open(sqlite.Open("test_tasks.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
}
