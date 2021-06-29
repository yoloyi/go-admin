package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
}

// GeneratePassword 生成密码
func (a *Auth) GeneratePassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PasswordVerify 密码验证
func (a *Auth) PasswordVerify(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func NewAuth() *Auth {
	return new(Auth)
}
