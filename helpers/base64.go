package helpers

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math"
)

// EncodeToBase64 encodes a string to Base64
func EncodeToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// DecodeFromBase64 decodes a Base64-encoded string
func DecodeFromBase64(s string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// EncodeUintToBase64 encodes a uint to a Base64 string
func EncodeUintToBase64(value uint) string {
	// Convert uint to a byte slice (4 bytes for uint32)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return base64.StdEncoding.EncodeToString(bytes)
}

// DecodeBase64ToUint decodes a Base64 string to a uint
func DecodeBase64ToUint(encoded string) (uint, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return 0, err
	}

	if len(decodedBytes) < 8 {
		return 0, fmt.Errorf("decoded data is too short")
	}

	value := binary.BigEndian.Uint64(decodedBytes)
	if value > math.MaxUint {
		return 0, fmt.Errorf("decoded value exceeds uint limit")
	}

	return uint(value), nil
}
