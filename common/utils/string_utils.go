package utils

import "strings"

func StringInclude(s string, sub string) bool {
	return strings.Contains(s, sub)
}
