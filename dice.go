package dice

import (
	"crypto/rand"
	"math/big"
)

// accepted characters
var (
	chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"
)

// RandInt
func RandInt(min, max int) int {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0
	}

	return min + int(r.Int64())
}

// RandString
func RandString(n int) string {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}

	return string(bytes)
}
