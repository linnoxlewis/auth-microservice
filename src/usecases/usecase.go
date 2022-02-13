package usecases

import (
	"auth-microservice/src/config"
	"auth-microservice/src/helpers"
	"auth-microservice/src/models"
	"auth-microservice/src/repository"
	"auth-microservice/src/services/jwt"
	"errors"
)

type UseCaseInterface interface {
	RegisterUser(email string, password string) (string, error)
	ConfirmRegister(token string) (*models.User, error)
	Login(email string, password string) (*jwt.Tokens, error)
}

var invalidUserDataErr = errors.New("Invalid email or password")
var userAlreadyExistErr = errors.New("User already exist")

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
	passwordHash, hashErr := helpers.GetPwdHash(password,u.envConf.GetPwdSalt())
	if hashErr != nil {
		return "", hashErr
	}
	claims := models.NewRegisterClaims(email, passwordHash, u.appConf.GetRegisterDuration())
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
		return nil, userAlreadyExistErr
	}
	user, errCreate := u.userRepo.CreateUser(tokenClaims.Email, tokenClaims.Password)
	if errCreate != nil {
		return nil, err
	}

	return user, nil
}

func (u *Usecase) Login(email string, password string) (*jwt.Tokens, error) {
	user := u.userRepo.GetUserByEmail(email)
	if !user.IsEmpty() {
		return nil, invalidUserDataErr
	}
	if !user.IsActive() {
		return nil, invalidUserDataErr
	}
	comparedPassword := helpers.ComparePassword(password,user.Password,u.envConf.GetPwdSalt())
	if comparedPassword != nil {
		return nil, invalidUserDataErr
	}

	accessClaims := models.NewAuthClaims(user.ID,u.appConf.GetAccessDuration())
	refreshClaims := models.NewAuthClaims(user.ID,u.appConf.GetAccessDuration())
	tokenAccess,tknAccErr := u.jwtService.GenerateToken(accessClaims,u.envConf.GetJwtAccessSecretKey())
	if tknAccErr != nil {
		return nil,tknAccErr
	}
	tokenRefresh,tknRefErr := u.jwtService.GenerateToken(refreshClaims,u.envConf.GetJwtRefreshSecretKey())
	if tknRefErr != nil {
		return nil,tknRefErr
	}

	return &jwt.Tokens{tokenAccess,tokenRefresh}, nil
}
