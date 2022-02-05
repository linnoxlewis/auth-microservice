package server

import (
	"auth-microservice/src/server/grpc/pb"
	"auth-microservice/src/usecases"
)
import "context"

type AuthServer struct {
	useCaseManager usecases.UseCaseInterface
	pb.UnimplementedAuthServer
}

func NewAuthServer(useCaseManager usecases.UseCaseInterface) * AuthServer {
	return &AuthServer{
		useCaseManager: useCaseManager,
	}
}

func (a *AuthServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	panic("implement me")
}

func (a *AuthServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	panic("implement me")
}

func (a *AuthServer) RecoverPassword(ctx context.Context, request *pb.RecoverPasswordRequest) (*pb.RecoverPasswordResponse, error) {
	panic("implement me")
}

func (a *AuthServer) Verify(ctx context.Context, request *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	panic("implement me")
}

func (a *AuthServer) UpdateTokens(ctx context.Context, request *pb.UpdateTokensRequest) (*pb.UpdateTokensResponse, error) {
	panic("implement me")
}


