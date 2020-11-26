package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthRouter(router *gin.RouterGroup) {
	r := router.Group("auth")
	r.GET("1", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"1": 1,
		})
	})
}
