package repository

import "auth-microservice/src/models"

type UserRepoInterface interface {
	GetUserByEmail(email string) (*models.User,error)
	GetUserById(id uint) (*models.User,error)
	CreateUser(email string, password string) (*models.User, error)
}
