// +build wireinject

package services

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
	"go-admin/internal/services/auth"
)

func NewAuthService() auth.Service {
	wire.Build(auth.ServiceWireSet, repositories.UserRepositoryWireSet)
	return auth.Service{}
}
