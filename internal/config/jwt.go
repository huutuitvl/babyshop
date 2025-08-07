package config

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtSecret = []byte("secret_key") // có thể để trong .env

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}
