package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kalamet/jdoc/model"
)

var jwtKey = []byte("ahlwerlsfdjiKflkeooo==-")
var method = jwt.SigningMethodHS256

type Claims struct {
	UserId int64
	jwt.MapClaims
}

func ReleaseToken(user model.User) (string, error) {
	//过期时间 7天
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		MapClaims: jwt.MapClaims{
			"exp": expirationTime.Unix(),
			"iss": "github.com/kalamet/jdoc",
			"sub": "user token",
			"iat": time.Now().Unix(),
		},
	}
	t := jwt.NewWithClaims(method, claims)

	tokenString, err := t.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(token string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return t, claims, err
}
