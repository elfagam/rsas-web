package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken membuat token JWT untuk user yang berhasil login
func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku 24 jam

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
