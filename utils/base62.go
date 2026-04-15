package utils

import (
	"strings"
)

const (
	charset       = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base          = int64(len(charset))
	shortCodeSize = 7
)

func Encode(num int64) string {
	if num == 0 {
		return strings.Repeat("0", shortCodeSize)
	}

	var encoded []byte

	for num > 0 {
		remainder := num % base
		encoded = append([]byte{charset[remainder]}, encoded...)
		num = num / base
	}

	// do the padding up-to shortCodeSize
	for len(encoded) < shortCodeSize {
		encoded = append([]byte{'0'}, encoded...)
	}

	return string(encoded)
}
