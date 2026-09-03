package fingerprint

import (
	"crypto/sha256"
	"encoding/hex"
)

func Fingerprint(body []byte) string {

	normalized := Normalize(body)
	// Calculate the SHA256  hash to complete response body
	hash := sha256.Sum256(normalized)

	// Convert the binary hash to a hexadecimal string and return it
	return hex.EncodeToString(hash[:])
}
