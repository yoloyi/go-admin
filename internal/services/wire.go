// +build wireinject

package services

import "github.com/google/wire"

func NewUserService() *user {
	wire.Build(userServiceSet)
	return &user{}
}
