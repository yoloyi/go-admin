package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	router.GET("/user/index", func(c *gin.Context) {})
}
