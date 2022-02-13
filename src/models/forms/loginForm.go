package forms

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type LoginForm struct {
	Email    string
	Password string
}

func NewLoginForm(email string, password string) *LoginForm {
	return &LoginForm{email, password}
}

func (l *LoginForm) Validate() error {
	return validation.ValidateStruct(l,
		validation.Field(&l.Email, validation.Required, is.Email),
		validation.Field(&l.Password, validation.Required))
}
