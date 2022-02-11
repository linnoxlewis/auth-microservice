package usecases

import (
	"auth-microservice/src/config"
	"auth-microservice/src/models"
	"auth-microservice/src/repository"
	"auth-microservice/src/services/jwt"
	"errors"
)

type UseCaseInterface interface {
	RegisterUser(email string, password string) (string, error)
	ConfirmRegister(token string) (*models.User, error)
}

type Usecase struct {
	appConf    *config.AppConf
	envConf    *config.EnvConfig
	jwtService jwt.JwtInterface
	userRepo   repository.UserRepoInterface
}

func NewUseCase(appConf *config.AppConf,
	envConf *config.EnvConfig,
	jwtSrv jwt.JwtInterface,
	usrRepo repository.UserRepoInterface) *Usecase {
	return &Usecase{
		appConf:    appConf,
		envConf:    envConf,
		jwtService: jwtSrv,
		userRepo:   usrRepo,
	}
}

func (u *Usecase) RegisterUser(email string, password string) (string, error) {
	claims := models.NewRegisterClaims(email, password, u.appConf.GetRegisterDuration())
	token, err := u.jwtService.GenerateToken(claims, u.envConf.GetJwtRegSecretKey())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *Usecase) ConfirmRegister(token string) (*models.User, error) {
	tokenClaims, err := u.jwtService.ParseRegisterToken(token, u.envConf.GetJwtRegSecretKey())
	if err != nil {
		return nil, err
	}
	userExist := u.userRepo.GetUserByEmail(tokenClaims.Email)
	if !userExist.IsEmpty() {
		return nil, errors.New("User already exist")
	}
	user, errCreate := u.userRepo.CreateUser(tokenClaims.Email, tokenClaims.Password)
	if errCreate != nil {
		return nil, err
	}

	return user, nil
}
