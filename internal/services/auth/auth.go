package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
	"go-admin/internal/routers/requests"
	auth2 "go-admin/internal/routers/requests/auth"
	authUtil "go-admin/internal/utils/auth"
)

var ServiceWireSet = wire.NewSet(NewAuth, authUtil.NewAuth)

type Service struct {
	userRepository *repositories.User
	authUtil       *authUtil.Auth
}

func NewAuth(userRepository *repositories.User, authUtil *authUtil.Auth) *Service {
	return &Service{
		userRepository: userRepository,
		authUtil:       authUtil,
	}
}

func (auth *Service) LoginService(ctx *gin.Context) {
	var loginRequest auth2.LoginRequest
	ctx.BindJSON(&loginRequest)
	if !requests.Validate(ctx, loginRequest) {
		return
	}
}
