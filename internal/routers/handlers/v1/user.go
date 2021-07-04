package v1

import (
	"github.com/gin-gonic/gin"
	"go-admin/internal/services"
)

type user struct {
}

// ChangePasswordHandler 用户修改密码
func (a user) ChangePasswordHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		s := services.NewUserService()
		s.UserChangePassword(context)
	}
}

// RegisterUserRouter 注册用户相关的路由
func RegisterUserRouter(router *gin.RouterGroup) {
	u := user{}
	r := router.Group("user")
	{
		r.PATCH("/changePassword", u.ChangePasswordHandler())
	}
}
