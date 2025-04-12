package random

import (
	"math/rand"
	"time"
)

func NewRandomAlias(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[rnd.Intn(len(charset))]
	}

	return string(b)
}
