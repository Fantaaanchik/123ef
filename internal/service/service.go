package service

import (
	"fmt"
	"repeatTestProject/models"
)

type Services struct {
	Repository RepositoryInterface
}

func NewService(repo RepositoryInterface) *Services {
	return &Services{Repository: repo}
}

type RepositoryInterface interface {
	GetUserFromDB() ([]models.User, error)
	AddNewUserToDB(users models.User) error
	DeleteUserByID(userID string) error
	UpdateUserDataFromDB(users *models.User) error
	GetUserByID(id string) (*models.User, error)
}

func (s Services) GetUserFromDB() ([]models.User, error) {
	return s.Repository.GetUserFromDB()
}

func (s Services) AddNewUserToDB(user models.User) error {
	return s.Repository.AddNewUserToDB(user)
}

func (s Services) UpdateUserDataFromDB(userID string, user models.User) error {
	existingUser, err := s.Repository.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("cannot get user by id %s: %s", userID, err.Error())
	}
	existingUser.Fio = user.Fio
	existingUser.Number = user.Number

	err = s.Repository.UpdateUserDataFromDB(existingUser)
	if err != nil {
		return fmt.Errorf("cannot update user data, err: %s", err.Error())
	}
	return nil
}

func (s Services) DeleteUserDataFromDB(userID string) error {
	err := s.Repository.DeleteUserByID(userID)
	if err != nil {
		return fmt.Errorf("cannot delete user by id %s: %s", userID, err.Error())
	}
	return nil
}
