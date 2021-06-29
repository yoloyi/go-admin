// +build wireinject

package repositories

import "github.com/google/wire"

func NewUserRepository() User {
	wire.Build(UserRepositoryWireSet)
	return User{}
}
