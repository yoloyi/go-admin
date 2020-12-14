package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Jwt struct {
	key string
}

type JwtClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func (this *Jwt) GenerateJwtToken(userId uint, expireDuration time.Duration) (string, error) {
	now := time.Now()
	nowTimeStamp := now.Unix()
	jwtClaim := JwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowTimeStamp,
			ExpiresAt: now.Add(expireDuration).Unix(),
			IssuedAt:  nowTimeStamp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	ss, err := token.SignedString([]byte(this.key))
	return ss, err
}

func (this *Jwt) ParseJwtToken(jwtToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, new(JwtClaims), func(token *jwt.Token) (interface{}, error) {
		return []byte(this.key), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func NewJwt(key string) *Jwt {
	return &Jwt{key: key}
}
