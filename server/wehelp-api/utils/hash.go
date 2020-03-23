package utils

import (
	"crypto/md5"
	"fmt"
)

// StringToMd5 :
func StringToMd5(input string) string {
	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", hash)
}

// ValidataMd5 :
func ValidataMd5(input, md5hash string) bool {
	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", hash) == md5hash
}
