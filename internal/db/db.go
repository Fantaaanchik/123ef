package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"repeatTestProject/config"
	"repeatTestProject/models"
)

var db *gorm.DB

func ConnectionToDB() *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(config.Configure.DB), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB, err: ", err.Error())
	}

	err = db.AutoMigrate(&models.User{})

	return db
}

func GetDB() *gorm.DB {
	return db
}
