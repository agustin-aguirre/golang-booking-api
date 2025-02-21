package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretlyconfidentialkey123$"

func GenerateToken(userId int64, email string) (string, error) {
	signingMethod := jwt.SigningMethodHS256
	claims := jwt.MapClaims{
		"email":      "",
		"userId":     "",
		"expiration": time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString([]byte(secretKey))
}
