package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"wheel/config"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(postgres.Open(config.GetConfig().DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	ApplyMigrations(db)

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
