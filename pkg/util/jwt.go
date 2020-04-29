package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"im/pkg/setting"
	"time"
)


var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
	jwt.StandardClaims
}


func GenerateToken(username, password string) (string, error) {
	now := time.Now()
	expireTime := now.Add(setting.TokenLife)

	claims := Claims{
		Username:       username,
		Password:       password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "thchat",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
