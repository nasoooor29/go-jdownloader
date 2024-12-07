package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// generateHMACSignature creates an HMAC-SHA256 signature.
func generateHMACSignature(data, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
