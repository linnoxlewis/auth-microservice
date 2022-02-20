package models

import (
	"auth-microservice/src/helpers"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Status   int
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

func (u *User) IsActive() bool {
	return u.Status == helpers.ACTIVE
}

func (u *User) IsBanned() bool {
	return u.Status == helpers.BANNED
}

func (u *User) IsNotActive() bool {
	return u.Status == helpers.NOT_ACTIVE
}
