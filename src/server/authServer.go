package server

import (
	"auth-microservice/src/models/forms"
	"auth-microservice/src/server/grpc/pb"
	"auth-microservice/src/usecases"
	"context"
)

type AuthServer struct {
	useCaseManager usecases.UseCaseInterface
	pb.UnimplementedAuthServer
}

func NewAuthServer(useCaseManager usecases.UseCaseInterface) *AuthServer {
	return &AuthServer{
		useCaseManager: useCaseManager,
	}
}

func (a *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.Token, error) {
	registerForm := forms.NewRegisterForm(request.GetEmail(), request.GetPassword())
	errValidate := registerForm.Validate()
	if errValidate != nil {
		return nil, errValidate
	}
	token, err := a.useCaseManager.RegisterUser(registerForm.Email, registerForm.Password)
	if err != nil {
		return nil, err
	}
	return &pb.Token{Token: token}, nil
}

func (a *AuthServer) ConfirmRegister(ctx context.Context, request *pb.Token) (*pb.User, error) {
	tokenForm := forms.NewTokenForm(request.GetToken())
	errValidate := tokenForm.Validate()
	if errValidate != nil {
		return nil, errValidate
	}

	confirmedUser, err := a.useCaseManager.ConfirmRegister(tokenForm.Token)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Email:  confirmedUser.Email,
		UserId: int64(confirmedUser.ID),
	}, nil
}

func (a *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.Tokens, error) {
	panic("implement me")
}
