package random

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateHex returns a "byte * 2" hex string
// after encoding "byte" random bytes in hex
func GenerateHex(bytes int) (string, error) {
	b, err := GenerateBytes(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
