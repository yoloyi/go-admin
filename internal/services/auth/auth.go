package auth

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
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

func (auth *Service) LoginService() {

}
