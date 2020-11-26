// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package repositories

import (
	"monitor/internal/models"
)

// Injectors from wire.go:

func NewUserRepository() *User {
	db := models.GetDB()
	user := NewUser(db)
	return user
}