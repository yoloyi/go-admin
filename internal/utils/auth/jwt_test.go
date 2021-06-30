package auth

import (
	"testing"
	"time"
)

func TestNewJwt(t *testing.T) {
	jwtUtil := NewJwt("123123123123")
	token, err := jwtUtil.GenerateUserJwtToken(123, 10*time.Hour)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestJwt_ParseJwtToken(t *testing.T) {
	jwtUtil := NewJwt("123123123123")
	token, err := jwtUtil.GenerateUserJwtToken(123, 1*time.Second)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
	claim, err := jwtUtil.ParseJwtToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claim)
}
