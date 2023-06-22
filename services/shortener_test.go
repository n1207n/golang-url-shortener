package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var userId = "a531d8fc-a391-498a-99c1-eb25ac62fe2c"

func TestShortener(t *testing.T) {
	originalUrl1 := "https://www.google.com"
	shortUrl1 := CreateShortUrl(originalUrl1, userId)

	originalUrl2 := "https://www.amazon.com"
	shortUrl2 := CreateShortUrl(originalUrl2, userId)

	originalUrl3 := "https://www.instagram.com"
	shortUrl3 := CreateShortUrl(originalUrl3, userId)

	assert.Equal(t, "ScvuNxB8NF5", shortUrl1)
	assert.Equal(t, "HcwMw9NwFZ1", shortUrl2)
	assert.Equal(t, "CrKjD9my6Eo", shortUrl3)
}
