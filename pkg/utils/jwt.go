package utils

import (
	"github.com/Kezume/BE-SPLITBILL.KUY/config"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte(config.AppConfig.JWTSecret)

func GenerateToken(id, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
	})
	return token.SignedString(SECRET_KEY)
}
