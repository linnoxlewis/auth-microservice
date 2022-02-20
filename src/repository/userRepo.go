package repository

import (
	"auth-microservice/src/helpers"
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

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := u.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	if err := u.db.First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) CreateUser(email string, password string) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Password: password,
		Status:   helpers.ACTIVE,
	}
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
