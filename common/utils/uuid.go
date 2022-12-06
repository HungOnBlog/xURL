package utils

import (
	"math/rand"

	"github.com/lithammer/shortuuid/v4"
)

func GenShortUUID() string {
	return shortuuid.New()
}

func GenFullUUID(namespace string) string {
	return shortuuid.NewWithNamespace(namespace)
}

// Generate random string with length n
func GenRandomString(length int) string {
	base62 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var result string
	for i := 0; i < length; i++ {
		result += string(base62[rand.Intn(len(base62))])
	}
	return result
}
