package util

import (
    "time"
    "github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("NHvbW`Yw3BgPec>=!{T-)D")

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Second * 30).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
