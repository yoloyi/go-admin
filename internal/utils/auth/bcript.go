package auth

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword 生成密码
func GeneratePassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// PasswordVerify 密码验证
func PasswordVerify(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Debugf("密码验证错误，%v", err)
		return false
	}
	return true
}
