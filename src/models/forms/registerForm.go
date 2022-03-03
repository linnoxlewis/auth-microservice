package forms

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"strings"
)

const MIN_PASSWORD_ENTROPY = 60

type RegisterForm struct {
	Email    string
	Password string
}

func NewRegisterForm(email string, password string) *RegisterForm {
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)
	return &RegisterForm{email, password}
}

func NewRegisterFormIngot() *RegisterForm {
	return &RegisterForm{}
}
func (r *RegisterForm) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.By(checkPassword)))
}

func checkPassword(value interface{}) error {
	return passwordvalidator.Validate(fmt.Sprintf("%v", value), MIN_PASSWORD_ENTROPY)
}
