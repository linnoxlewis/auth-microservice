package models

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type RegisterClaims struct {
	jwt.StandardClaims
	Email    string
	Password string
}

type AuthClaims struct {
	jwt.StandardClaims
	Uid uint
}

func NewRegisterClaims(email string, password string, duration time.Duration) *RegisterClaims {
	return &RegisterClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * duration).Unix(),
		},
		Email:    email,
		Password: password,
	}
}

func NewAuthClaims(userId uint, duration time.Duration) *AuthClaims {
	return &AuthClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * duration).Unix(),
		},
		Uid: userId,
	}
}
