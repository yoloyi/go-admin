// +build wireinject

package services

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
)

func NewAuthService() *Auth {
	wire.Build(AuthServiceWireSet, repositories.UserRepositoryWireSet)
	return &Auth{}
}
