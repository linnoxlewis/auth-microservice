package usecases

import (
	"auth-microservice/src/repository"
	"auth-microservice/src/services/jwt"
)

type UseCaseInterface interface {

}

type Usecase struct {
	jwtService jwt.JwtInterface
	userRepo repository.UserRepoInterface
}

func NewUseCase(jwtSrv jwt.JwtInterface,usrRepo repository.UserRepoInterface) *Usecase {
	return &Usecase{
		jwtService: jwtSrv,
		userRepo: usrRepo,
	}
}
