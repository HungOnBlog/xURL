package utils

// Return hash id from int
func HashId(id uint) string {
	base62 := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var hash string
	for id > 0 {
		hash = string(base62[id%62]) + hash
		id = id / 62
	}
	return hash
}
