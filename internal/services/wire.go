// +build wireinject

package services

import (
	"github.com/google/wire"
	"monitor/internal/models/repositories"
)

func NewAuthService() *auth {
	wire.Build(AuthServiceWireSet, repositories.UserRepositoryWireSet)
	return &auth{}
}
