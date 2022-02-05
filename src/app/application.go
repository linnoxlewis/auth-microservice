package app

import (
	"auth-microservice/src/config"
	"auth-microservice/src/repository"
	"auth-microservice/src/server"
	"auth-microservice/src/services/db"
	"auth-microservice/src/services/jwt"
	"auth-microservice/src/usecases"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) {
	config.Init()

	database :=db.GetDB()
	defer db.CloseDB(database)

	userRepo := repository.NewUserRepository(database)
	jwtService :=jwt.NewJwtService()

	useCaseManager := usecases.NewUseCase(jwtService,userRepo)

	srv := server.NewGrpcServer(":80",useCaseManager)
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
