package repositories

import (
	"github.com/google/wire"
	"go-admin/internal/models"
	"go-admin/internal/models/entities"
	"gorm.io/gorm"
)

var UserRepositoryWireSet = wire.NewSet(NewUser, models.GetDB)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{
		DB: db,
	}
}

// Create 创建用户
func (u User) Create(users entities.User) error {
	return u.DB.Create(users).Error
}

func (u User) GetUserByUserName(username string) (entities.User, error) {
	user := entities.User{}
	err := u.DB.Where("username = ?", username).First(&user).Error

	return user, err
}

func (u User) First(conds ...interface{}) (entities.User, error) {
	user := entities.User{}
	err := u.DB.First(&user, conds...).Error
	return user, err
}


