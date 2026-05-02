package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Kezume/BE-SPLITBILL.KUY/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMidleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization is missing",
			})
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Invalid Signing Method")
			}

			return config.AppConfig.JWTSecret, nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {
			userID := claims["id"].(string)
			c.Set("id", userID)
		}

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})
			return
		}

		c.Next()
	}
}
