package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var userId = uuid.New().String()

func TestShortener(t *testing.T) {
	originalUrl1 := "https://www.google.com"
	shortUrl1 := CreateShortUrl(originalUrl1, userId)

	originalUrl2 := "https://www.amazon.com"
	shortUrl2 := CreateShortUrl(originalUrl2, userId)

	originalUrl3 := "https://www.instagram.com"
	shortUrl3 := CreateShortUrl(originalUrl3, userId)

	assert.Equal(t, shortUrl1, "hB5TnJXinCG")
	assert.Equal(t, shortUrl2, "AccyLG2zZFu")
	assert.Equal(t, shortUrl3, "5ip8weqia8n")
}
