package handlers

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/services"
)

type auth struct {
}

func (a *auth) loginHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		authService := services.NewAuthService()
		authService.LoginService(context)
	}
}

func RegisterAuthRouter(router *gin.RouterGroup) {
	var auth auth
	r := router.Group("auth")
	{
		r.POST("login", auth.loginHandler())
		r.GET("login", func(context *gin.Context) {
			
		})
	}
}
