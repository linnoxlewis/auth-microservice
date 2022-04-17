package usecases

import (
	"auth-microservice/src/config"
	"auth-microservice/src/helpers"
	"auth-microservice/src/log"
	"auth-microservice/src/models"
	"auth-microservice/src/models/consts"
	"auth-microservice/src/repository"
	"auth-microservice/src/services/jwt"
	"errors"
)

type UseCaseInterface interface {
	RegisterUser(email string, password string) (string, error)
	ConfirmRegister(token string) (*models.User, error)
	Login(email string, password string) (*jwt.Tokens, error)
	GetTokensByRefresh(refreshToken string) (*jwt.Tokens, error)
	Verify(accessToken string) (bool, error)
}

var (
	errInvalidUserData  = errors.New(consts.ErrInvalidEmailOrPassword)
	errUserAlreadyExist = errors.New(consts.ErrUserAlreadyExist)
	errUserNotFound     = errors.New(consts.ErrUserNotFound)
	errInternalServer   = errors.New(consts.ErrInternalServer)
	errUserStatus       = errors.New(consts.ErrUserNotActive)
	errTokenInvalid     = errors.New(consts.ErrInvalidToken)
)

type Usecase struct {
	appConf    *config.AppConf
	envConf    *config.EnvConfig
	jwtService jwt.JwtInterface
	userRepo   repository.UserRepoInterface
	logger     *log.Logger
}

func NewUseCase(appConf *config.AppConf,
	envConf *config.EnvConfig,
	jwtSrv jwt.JwtInterface,
	usrRepo repository.UserRepoInterface,
	logger *log.Logger) *Usecase {
	return &Usecase{
		appConf:    appConf,
		envConf:    envConf,
		jwtService: jwtSrv,
		userRepo:   usrRepo,
		logger:     logger,
	}
}

func (u *Usecase) RegisterUser(email string, password string) (string, error) {
	existUser := u.userRepo.GetUserByEmail(email)
	if !existUser.IsEmpty() {
		return "", errUserAlreadyExist
	}

	passwordHash, err := helpers.GetPwdHash(password, u.envConf.GetPwdSalt())
	if err != nil {
		u.logger.ErrorLog.Println("hash password error: ", err)

		return "", errInternalServer
	}
	claims := models.NewRegisterClaims(email, passwordHash, u.appConf.GetRegisterDuration())
	token, err := u.jwtService.GenerateToken(claims, u.envConf.GetJwtRegSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("generate token error: ", err)

		return "", errInternalServer
	}

	return token, nil
}

func (u *Usecase) ConfirmRegister(token string) (*models.User, error) {
	tokenClaims, err := u.jwtService.ParseRegisterToken(token, u.envConf.GetJwtRegSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("parse token error: ", err)

		return nil, errTokenInvalid
	}

	existUser := u.userRepo.GetUserByEmail(tokenClaims.Email)
	if !existUser.IsEmpty() {
		return nil, errUserAlreadyExist
	}

	user, err := u.userRepo.CreateUser(tokenClaims.Email, tokenClaims.Password)
	if err != nil {
		u.logger.ErrorLog.Println("create user error: ", err)

		return nil, errInternalServer
	}

	return user, nil
}

func (u *Usecase) Login(email string, password string) (*jwt.Tokens, error) {
	user := u.userRepo.GetUserByEmail(email)
	if user.IsEmpty() {
		return nil, errInvalidUserData
	}
	if !user.IsActive() {
		return nil, errUserStatus
	}

	comparedPassword := helpers.ComparePassword(password, user.Password, u.envConf.GetPwdSalt())
	if comparedPassword != nil {
		return nil, errInvalidUserData
	}

	return u.getTokens(user.ID)
}

func (u *Usecase) GetTokensByRefresh(refreshToken string) (*jwt.Tokens, error) {
	tokenClaims, err := u.jwtService.ParseAuthToken(refreshToken, u.envConf.GetJwtRefreshSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("parse token error: ", err)

		return nil, errTokenInvalid
	}

	user := u.userRepo.GetUserById(tokenClaims.Uid)
	if user.IsEmpty() {
		return nil, errUserNotFound
	}

	return u.getTokens(user.ID)
}

func (u *Usecase) Verify(accessToken string) (bool, error) {
	verify, err := u.jwtService.Verify(accessToken, u.envConf.GetJwtAccessSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("verify access token error: ", err)

		return false, errTokenInvalid
	}

	return verify, nil
}

func (u *Usecase) getTokens(userId uint) (*jwt.Tokens, error) {
	accessClaims := models.NewAuthClaims(userId, u.appConf.GetAccessDuration())
	refreshClaims := models.NewAuthClaims(userId, u.appConf.GetAccessDuration())

	tokenAccess, err := u.jwtService.GenerateToken(accessClaims, u.envConf.GetJwtAccessSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("generate access token error: ", err)

		return nil, errInternalServer
	}

	tokenRefresh, err := u.jwtService.GenerateToken(refreshClaims, u.envConf.GetJwtRefreshSecretKey())
	if err != nil {
		u.logger.ErrorLog.Println("generate refresh token error: ", err)

		return nil, errInternalServer
	}

	return &jwt.Tokens{
		AccessToken:  tokenAccess,
		RefreshToken: tokenRefresh,
	}, nil
}
