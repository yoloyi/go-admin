package services

import (
	"github.com/google/wire"
	"monitor/internal/models/repositories"
)

var AuthServiceWireSet = wire.NewSet(newAuthService)

type auth struct {
	userRepository *repositories.User
}

func newAuthService(userRepository *repositories.User) *auth {
	return &auth{userRepository: userRepository}
}
