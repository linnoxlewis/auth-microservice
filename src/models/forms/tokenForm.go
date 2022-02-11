package forms

import validation "github.com/go-ozzo/ozzo-validation/v4"

type TokenFrom struct {
	Token string
}

func NewTokenForm(token string) *TokenFrom {
	return &TokenFrom{token}
}

func (t *TokenFrom) Validate() error {
	return validation.ValidateStruct(t, validation.Field(&t.Token, validation.Required))
}
