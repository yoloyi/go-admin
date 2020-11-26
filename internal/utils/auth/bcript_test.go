package auth

import (
	"strconv"
	"testing"
)

func TestBcript(t *testing.T) {

	password := "password123454657645322"
	hashPassword, err := GeneratePassword(password)
	if err != nil {
		t.Fatal(err)
	}
	if !PasswordVerify(password, hashPassword) {
		t.Fatal("Hash 解析错误")
	}
	t.Logf("password: %s, hash password: %s. TEST SUCCESS", password, hashPassword)
}

func BenchmarkBcript(b *testing.B) {
	for i := 0; i < b.N; i++ {
		password := strconv.Itoa(i)
		hashPassword, err := GeneratePassword(password)
		if err != nil {
			b.Fatal(err)
		}
		if !PasswordVerify(password, hashPassword) {
			b.Fatal("Hash 解析错误")
		}
	}
}
