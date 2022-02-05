package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type JwtInterface interface {
	GenerateToken(claims jwt.Claims) (string, error)
	ParseToken(strToken string) (*jwt.MapClaims, error)
}
type JwtService struct {
	key      string
	issue    string
	audience string
}

func NewJwtService() *JwtService {
	secretKey := fmt.Sprintf("%s", viper.Get("jwt.secretKey"))
	issue := fmt.Sprintf("%s", viper.Get("jwt.issue"))
	audience := fmt.Sprintf("%s", viper.Get("jwt.audience"))
	return &JwtService{
		key:      secretKey,
		issue:    issue,
		audience: audience,
	}
}

func (j *JwtService) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.key))
}

func (j *JwtService) ParseToken(strToken string) (*jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}
	token, parseErr := jwt.ParseWithClaims(strToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid method signning!")
		}
		return []byte(j.key), nil
	})
	if parseErr != nil {
		return nil, parseErr
	}
	if !token.Valid {
		return nil, errors.New("Invalid token error")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Get token claims error")
	}
}
