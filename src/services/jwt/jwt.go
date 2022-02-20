package jwt

import (
	"auth-microservice/src/models"
	"errors"
	"github.com/golang-jwt/jwt"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var invalidSingErr = errors.New("Invalid method sign!")
var claimsError = errors.New("Get token claims error")

type JwtService struct {}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (j *JwtService) GenerateToken(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func (j *JwtService) ParseRegisterToken(strToken string, secretKey string) (*models.RegisterClaims, error) {
	claims := &models.RegisterClaims{}
	token, err := j.parse(strToken, claims, secretKey)
	if err != nil {
		return nil, err
	}
	clms, ok := token.Claims.(*models.RegisterClaims)
	if ok && token.Valid {
		return clms, nil
	} else {
		return nil, claimsError
	}
}

func (j *JwtService) ParseAuthToken(strToken string, secretKey string) (*models.AuthClaims, error) {
	claims := &models.AuthClaims{}
	token, err := j.parse(strToken, claims, secretKey)
	if err != nil {
		return nil, err
	}
	clms, ok := token.Claims.(*models.AuthClaims)
	if ok && token.Valid {
		return clms, nil
	} else {
		return nil, claimsError
	}
}

func (j *JwtService) Verify(strToken string, secretKey string) (bool, error) {
	token, parseErr := j.parse(strToken, jwt.MapClaims{}, secretKey)
	if parseErr != nil {
		return false, parseErr
	}

	return token.Valid, nil
}

func (j *JwtService) parse(strToken string, model jwt.Claims, key string) (*jwt.Token, error) {
	token, parseErr := jwt.ParseWithClaims(strToken, model, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidSingErr
		}
		return []byte(key), nil
	})

	return token, parseErr
}
