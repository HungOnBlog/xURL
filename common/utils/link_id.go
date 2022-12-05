package utils

// EncodeLinkID encodes an integer to a string
func LinkId(id uint) string {
	base32 := "0123456789bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

	var linkId string
	for id > 0 {
		linkId = string(base32[id%32]) + linkId
		id /= 32
	}

	return linkId
}
