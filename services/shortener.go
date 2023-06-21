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

func createShortUrl(originalLink string, userId string) string {
	urlBytes := getSha256(originalLink)
	hashedNumber := new(big.Int).SetBytes(urlBytes).Uint64()

	encoded := getBase58([]byte(fmt.Sprintf("%d", hashedNumber)))
	return encoded
}
