package app

import (
	"auth-microservice/src/config"
	"auth-microservice/src/server"
	"auth-microservice/src/services/db"
	"context"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) {
	config.Init()

	database :=db.GetDB()
	defer db.CloseDB(database)

	srv := server.NewGrpcServer(":80")
	go srv.StartServer()
	defer srv.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
