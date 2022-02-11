package jwt

import (
	"auth-microservice/src/config"
	"auth-microservice/src/models"
	"errors"
	"github.com/golang-jwt/jwt"
)

type JwtInterface interface {
	GenerateToken(claims jwt.Claims, secretKey string) (string, error)
	ParseRegisterToken(strToken string, secretKey string) (*models.RegisterClaims, error)
	ParseAuthToken(strToken string, secretKey string) (*models.AuthClaims, error)
	Verify(strToken string, secretKey string) (bool, error)
}

type JwtService struct {
	issue    string
	audience string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

var invalidSingErr = errors.New("Invalid method sign!")
var claimsError = errors.New("Get token claims error")

func NewJwtService(appCfg *config.AppConf) *JwtService {
	return &JwtService{
		issue:    appCfg.GetJwtIssue(),
		audience: appCfg.GetJwtAudience(),
	}
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
	token, parseErr := j.parse(strToken, jwt.StandardClaims{}, secretKey)
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
