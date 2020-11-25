package services

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"monitor/internal/models"
)

type user struct {
	db *gorm.DB
}

func newUser(db *gorm.DB) *user {
	return &user{
		db: db,
	}
}

var userServiceSet = wire.NewSet(newUser, models.GetDB)
