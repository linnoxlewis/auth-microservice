package jwt

import (
	"auth-microservice/src/models"
	"github.com/golang-jwt/jwt"
)

type JwtInterface interface {
	GenerateToken(claims jwt.Claims, secretKey string) (string, error)
	ParseRegisterToken(strToken string, secretKey string) (*models.RegisterClaims, error)
	ParseAuthToken(strToken string, secretKey string) (*models.AuthClaims, error)
	Verify(strToken string, secretKey string) (bool, error)
}
