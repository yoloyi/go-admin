// +build wireinject

package services

import (
	"github.com/google/wire"
	"go-admin/internal/models/repositories"
)

func NewAuthService() *auth {
	wire.Build(AuthServiceWireSet, repositories.UserRepositoryWireSet)
	return &auth{}
}
