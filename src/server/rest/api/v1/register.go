package v1

import (
	"auth-microservice/src/log"
	"auth-microservice/src/server/rest/hundler/v1"
	"auth-microservice/src/usecases"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router *gin.Engine, us usecases.UseCaseInterface, logger *log.Logger) {
	grp := router.Group("/api/v1/auth")
	hdr := v1.NewHundler(us, logger)

	grp.POST("/register", hdr.Register)
	grp.GET("/confirm-register", hdr.ConfirmRegister)
	grp.POST("/login", hdr.Login)
	grp.POST("/verify", hdr.Verify)
	grp.POST("/update-tokens", hdr.UpdateTokens)
}
