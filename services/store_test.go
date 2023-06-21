package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStorageService *StorageService

func init() {
	testStorageService = NewStorageService()
}

func TestStorageInit(t *testing.T) {
	assert.True(t, testStorageService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	originalUrl := "https://www.google.com"
	userId := uuid.New().String()
	shortUrl := "fT34gvbscdDF"

	SaveUrlMapping(shortUrl, originalUrl, userId)

	retrievedUrl := GetUrlMapping(shortUrl)

	assert.Equal(t, originalUrl, retrievedUrl)
}
