package utils

import "github.com/lithammer/shortuuid/v4"

func GenShortUUID() string {
	return shortuuid.New()
}
