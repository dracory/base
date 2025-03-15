package str

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 converts a string to MD5 hash
func MD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
