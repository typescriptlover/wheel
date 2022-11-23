package db

import (
	"wheel/models"

	"gorm.io/gorm"
)

func ApplyMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
