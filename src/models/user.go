package models

import (
	"auth-microservice/src/helpers/userStatus"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Status   int
}

func (u *User) IsActive() bool {
	return u.Status == userStatus.ACTIVE
}

func (u *User) IsBanned() bool {
	return u.Status == userStatus.BANNED
}

func (u *User) IsNotActive() bool {
	return u.Status == userStatus.NOT_ACTIVE
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}
