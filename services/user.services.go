package services

import "go-curd-demo/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpateUser(*models.User) error
	DeleteUser(*models.User) error
}
