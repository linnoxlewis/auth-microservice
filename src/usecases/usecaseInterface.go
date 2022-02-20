package usecases

import (
	"auth-microservice/src/models"
	"auth-microservice/src/services/jwt"
)

type UseCaseInterface interface {
	RegisterUser(email string, password string) (string, error)
	ConfirmRegister(token string) (*models.User, error)
	Login(email string, password string) (*jwt.Tokens, error)
	GetTokensByRefresh(refreshToken string) (*jwt.Tokens, error)
	Verify(accessToken string) (bool,error)
}
