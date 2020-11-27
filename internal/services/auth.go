package services

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
)

var AuthServiceWireSet = wire.NewSet(NewAuth)

type Auth struct {
	userRepository *repositories.User
}

func NewAuth(userRepository *repositories.User) *Auth {
	return &Auth{userRepository: userRepository}
}
