package repository

import (
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
	user := models.NewUserIngot()
	u.db.Where("email = ?", email).First(user)

	return user
}

func (u *UserRepository) GetUserById(id uint) *models.User {
	user := models.NewUserIngot()
	u.db.Where("id = ?", id).First(user)

	return user
}

func (u *UserRepository) CreateUser(email string, password string) (*models.User, error) {
	user := &models.User{
		Email:    email,
		Password: password,
		Status:   models.UserStatusActive,
	}
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
