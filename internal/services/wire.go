// +build wireinject

package services

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
	"go-admin/internal/services/auth"
	"go-admin/internal/services/user"
)

func NewAuthService() auth.Service {
	wire.Build(auth.ServiceWireSet, repositories.UserRepositoryWireSet)
	return auth.Service{}
}

func NewUserService() user.Service {
	wire.Build(user.ServiceWireSet, repositories.UserRepositoryWireSet)
	return user.Service{}
}
