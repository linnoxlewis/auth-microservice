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
	GetTokensByRefresh(refreshToken string) (*jwt.Tokens, error)
	Verify(accessToken string) (bool,error)
}

var (
	invalidUserDataErr = errors.New("Invalid email or password")
	userAlreadyExistErr = errors.New("User already exist")
	userNotFound = errors.New("User not found")
)

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
	passwordHash, err := helpers.GetPwdHash(password,u.envConf.GetPwdSalt())
	if err != nil {
		return "", err
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
	user, err := u.userRepo.CreateUser(tokenClaims.Email, tokenClaims.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Usecase) Login(email string, password string) (*jwt.Tokens, error) {
	user := u.userRepo.GetUserByEmail(email)
	if user.IsEmpty() {
		return nil, invalidUserDataErr
	}
	if !user.IsActive() {
		return nil, invalidUserDataErr
	}

	comparedPassword := helpers.ComparePassword(password,user.Password,u.envConf.GetPwdSalt())
	if comparedPassword != nil {
		return nil, invalidUserDataErr
	}

	return u.getTokens(user.ID)
}

func (u *Usecase) GetTokensByRefresh(refreshToken string) (*jwt.Tokens, error) {
	tokenClaims, err := u.jwtService.ParseAuthToken(refreshToken, u.envConf.GetJwtRefreshSecretKey())
	if err != nil {
		return nil, err
	}
	user := u.userRepo.GetUserById(tokenClaims.Uid)
	if user.IsEmpty() {
		return nil, userNotFound
	}

	return u.getTokens(user.ID)
}

func (u *Usecase) Verify(accessToken string) (bool,error) {
	return u.jwtService.Verify(accessToken,u.envConf.GetJwtAccessSecretKey())
}

func (u *Usecase) getTokens(userId uint) (*jwt.Tokens, error) {
	accessClaims := models.NewAuthClaims(userId,u.appConf.GetAccessDuration())
	refreshClaims := models.NewAuthClaims(userId,u.appConf.GetAccessDuration())
	tokenAccess,err := u.jwtService.GenerateToken(accessClaims,u.envConf.GetJwtAccessSecretKey())
	if err != nil {
		return nil,err
	}
	tokenRefresh,err := u.jwtService.GenerateToken(refreshClaims,u.envConf.GetJwtRefreshSecretKey())
	if err != nil {
		return nil,err
	}

	return &jwt.Tokens{tokenAccess,tokenRefresh}, nil
}