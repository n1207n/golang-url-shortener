package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStorageService *StorageService

func init() {
	testStorageService = NewStorageService("localhost:6379", "", 0)
}

func TestStorageInit(t *testing.T) {
	assert.True(t, testStorageService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	originalUrl := "https://www.google.com"
	shortUrl := "fT34gvbscdDF"

	SaveUrlMapping(shortUrl, originalUrl)

	retrievedUrl := GetUrlMapping(shortUrl)

	assert.Equal(t, originalUrl, retrievedUrl)
}
