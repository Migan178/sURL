package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandomString(length int) string {
	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(buf)
}
