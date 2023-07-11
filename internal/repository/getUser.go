package repository

import (
	"gorm.io/gorm"
	"log"
	"repeatTestProject/internal/db"
	"repeatTestProject/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

type UserRepository struct {
	UserRep models.User
}

func (r *Repository) GetUserFromDB() ([]models.User, error) {
	var users []models.User

	err := db.GetDB().Find(&users).Error
	if err != nil {
		db.GetDB().Rollback()
		log.Println("Cannot get user from DB, err: ", err.Error())
	}
	return users, nil
}
