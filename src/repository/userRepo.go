package repository

import (
	"auth-microservice/src/helpers/userStatus"
	"auth-microservice/src/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByEmail(email string) *models.User {
	user := &models.User{}
	u.db.Where("email = ?", email).First(user)

	return user
}

func (u *UserRepository) GetUserById(id int) *models.User {
	user := &models.User{}
	u.db.First(user, id)

	return user
}

func (u *UserRepository) CreateUser(email string, password string) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Password: password,
		Status:   userStatus.ACTIVE,
	}
	err := u.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
