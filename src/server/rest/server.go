package rest

import (
	"auth-microservice/src/log"
	"auth-microservice/src/server/rest/api/v1"
	"auth-microservice/src/usecases"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpSrv *http.Server
	logger  *log.Logger
}

type RestServerInterface interface {
	StartServer()
	StopServer()
}

func NewServer(usecases usecases.UseCaseInterface,logger *log.Logger, port string) *Server {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.RegisterEndpoints(router,usecases,logger)

	srv := &http.Server{
		Addr:     port,
		Handler:  router,
		ErrorLog: logger.ErrorLog,
	}

	return &Server{
		httpSrv: srv,
		logger:  logger,
	}
}

func (s *Server) StartServer() {
	if err := s.httpSrv.ListenAndServe(); err != nil {
		s.logger.ErrorLog.Fatalf("Failed to listen and serve: %+v", err)
	}
	s.logger.InfoLog.Println("Start REST server...")
}

func (s *Server) StopServer() {
	if err := s.httpSrv.Shutdown(context.Background()); err != nil {
		s.logger.ErrorLog.Fatalf("Failed stopped serve: %+v", err)
	}
	s.logger.InfoLog.Println("Shutting down server...")
}
