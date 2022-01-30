package server

import (
	"auth-microservice/src/server/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServerInterface interface {
	StartServer()
	StopServer()
}

type GrpcServer struct {
	server     *grpc.Server
	Port       string
}

func NewGrpcServer(port string) *GrpcServer {
	srv := grpc.NewServer()
	authSrv := NewAuthServer()
	pb.RegisterAuthServer(srv,authSrv)

	return &GrpcServer{server: srv, Port: port}
}

func (g *GrpcServer) StartServer() {
	log.Println("Auth server starting...")

	l,err := net.Listen("tcp",g.Port)
	if err != nil {
		panic(err)
	}
	err = g.server.Serve(l)
	if err != nil {
		panic(err)
	}
}

func (g *GrpcServer) StopServer() {
	log.Println("Auth server stopping...")
	g.server.Stop()
}
