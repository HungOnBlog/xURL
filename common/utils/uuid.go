package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

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

// Return MD5 hash of input string and current time
func GenApikey(input string) string {
	var valueFrom string
	if input == "" {
		valueFrom = GenRandomString(32)
	} else {
		valueFrom = input
	}

	return GenMD5Hash(valueFrom + time.Now().String())
}

// Generate MD5 hash of input string
func GenMD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
