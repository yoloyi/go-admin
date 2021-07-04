package entities

import (
	"go-admin/internal/utils/auth"
)

// Users 后台用户表
type User struct {
	Username string `gorm:"unique;index:udx_username;column:username;type:varchar(100);not null"` // 用户名
	Password string `gorm:"column:password;type:varchar(100);not null"`                           // 密码
	Status   int8   `gorm:"column:status;type:tinyint(1);not null"`                               // 状态
	BaseEntity
}

const (
	UserStatusNormal = 0 // 正常用户
	UserForbidden    = 1 // 禁用
)

// CheckUserNormal 判断用户是否为正常用户
func (u User) CheckUserNormal() bool {
	return u.Status == UserStatusNormal
}

// CheckPasswordValid 检查用户密码是否有效
func (u User) CheckPasswordValid(password string) bool {
	authUtils := auth.NewAuth()
	return authUtils.PasswordVerify(password, u.Password)
}

// GenerateNewPassword 生成新的密码，但不保存
func (u *User) GenerateNewPassword(password string) error {

	authUtils := auth.NewAuth()
	passwordHash, err := authUtils.GeneratePassword(password)

	if err != nil {
		return err
	}

	u.Password = passwordHash
	return nil
}
