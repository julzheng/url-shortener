package utils

import (
	"encoding/hex"
	"math/rand"
)

func RandSeq(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
