package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK,
		gin.H{
			"data":  data,
			"error": nil,
		})
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"data":  nil,
			"error": err.Error(),
		})
}
