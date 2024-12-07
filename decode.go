package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

func decrypt(secretToken, data string) (string, error) {
	// Split the secretToken into initVector and key
	if len(secretToken)%2 != 0 {
		return "", errors.New("secretToken length must be even")
	}
	mid := len(secretToken) / 2
	initVector := []byte(secretToken[:mid])
	key := []byte(secretToken[mid:])

	// Decode the base64 encoded data
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	// Create AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Ensure data length is a multiple of the block size
	if len(decodedData)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	// Decrypt data using CBC mode
	mode := cipher.NewCBCDecrypter(block, initVector)
	decryptedData := make([]byte, len(decodedData))
	mode.CryptBlocks(decryptedData, decodedData)

	// Remove padding
	decryptedData, err = unpad(decryptedData)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}

// unpad removes PKCS#7 padding
func unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is empty")
	}

	padding := int(data[length-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, errors.New("invalid padding size")
	}

	return data[:length-padding], nil
}

func secretCreate(email, password, domain string) []byte {
	// Convert email and domain to lowercase
	emailLower := strings.ToLower(email)
	domainLower := strings.ToLower(domain)

	// Create the SHA-256 hash
	hash := sha256.New()
	hash.Write([]byte(emailLower))
	hash.Write([]byte(password))
	hash.Write([]byte(domainLower))

	// Return the digest
	return hash.Sum(nil)
}
