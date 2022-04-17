package models

import (
	"github.com/jinzhu/gorm"
)

const UserStatusNotActive int = 1

const UserStatusActive int = 2

type User struct {
	gorm.Model
	Email    string
	Password string
	Status   int
}

func NewUserIngot() *User {
	return &User{}
}

func (u *User) IsEmpty() bool {
	return u.ID == 0
}

func (u *User) IsActive() bool {
	return u.Status == UserStatusActive
}

func (u *User) IsNotActive() bool {
	return u.Status == UserStatusNotActive
}
