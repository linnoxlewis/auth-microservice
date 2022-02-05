package repository

import "auth-microservice/src/models"

type UserRepoInterface interface {
	GetUserByEmail(email string) *models.User
	GetUserById(id int) *models.User
	CreateUser(email string, password string) (*models.User, error)
}
