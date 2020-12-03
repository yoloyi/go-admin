package repositories

import (
	"github.com/google/wire"
	"go-admin/internal/models"
	"go-admin/internal/models/entities"
	"gorm.io/gorm"
)

var UserRepositoryWireSet = wire.NewSet(NewUser, models.GetDB)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(users *entities.User) error {
	return u.db.Create(users).Error
}

func (u *User) GetUserByUserName(username string) (*entities.User, error) {
	user := &entities.User{}
	err := u.db.Where("username = ?", username).First(user).Error

	return user, err
}
