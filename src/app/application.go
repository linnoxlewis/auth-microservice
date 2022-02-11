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
	appCfg := config.Init()
	envCfg := config.NewEnvConfig()

	database := db.GetDB(envCfg)
	defer db.CloseDB(database)

	userRepo := repository.NewUserRepository(database)
	jwtService := jwt.NewJwtService(appCfg)
	useCaseManager := usecases.NewUseCase(appCfg, envCfg, jwtService, userRepo)

	srv := server.NewGrpcServer(envCfg.GetServerPort(), useCaseManager)
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
