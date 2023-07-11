package repository

import (
	"log"
	"repeatTestProject/internal/db"
	"repeatTestProject/models"
)

func (r *Repository) AddNewUserToDB(users models.User) error {
	err := db.GetDB().Create(&users).Error
	if err != nil {
		log.Println("Cannot add new User to DB, err: ", err.Error())
		return err
	}
	return nil
}
