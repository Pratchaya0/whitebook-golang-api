package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

// Hash generates a SHA-256 hash for the given string
func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
