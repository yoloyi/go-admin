package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
}

func (a *auth) loginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func RegisterAuthRouter(router *gin.RouterGroup) {
	var auth auth
	r := router.Group("auth")
	{
		r.GET("login", auth.loginHandler())
	}
}
