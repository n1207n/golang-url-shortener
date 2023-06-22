package services

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
)

// getSha256 returns raw bytes of SHA-256 hashed content
func getSha256(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

// getBase58 converts a byte array into string with Base58 encoding
func getBase58(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

// CreateShortUrl creates a short ID by hashing the original link then encoding it with base58
func CreateShortUrl(originalLink string, userId string) string {
	urlBytes := getSha256(fmt.Sprintf("%s|%s", originalLink, userId))

	// Generate a random bigint number as a vessel and set the value by SetBytes
	hashedNumber := new(big.Int).SetBytes(urlBytes).Uint64()

	hashedNumberBytes := []byte(fmt.Sprintf("%d", hashedNumber))
	encoded := getBase58(hashedNumberBytes)
	return encoded
}
