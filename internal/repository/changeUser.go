package repository

import (
	"log"
	"repeatTestProject/internal/db"
	"repeatTestProject/models"
)

func (r *Repository) UpdateUserDataFromDB(users *models.User) error {
	err := db.GetDB().Updates(&users).Error
	if err != nil {
		db.GetDB().Rollback()
		log.Println("Cannot update user data, err: ", err.Error())
		return err
	}
	return nil
}

func (r *Repository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := db.GetDB().Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println("Cannot get user by id, err: ", err.Error())
	}
	return &user, nil
}

func (r *Repository) DeleteUserByID(id string) error {
	err := db.GetDB().Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		log.Println("Cannot delete user by his id, err: ", err.Error())
		return err
	}
	return nil
}
