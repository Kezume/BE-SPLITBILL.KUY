package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomInviteCode() string {
	rand.Seed(time.Now().UnixNano())

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]rune, 8)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
