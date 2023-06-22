package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/n1207n/golang-url-shortener/services"
	"net/http"
)

type UrlCreateRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
}

// GenerateShortUrl generates the shortened url from original url, persists the record, and return the response
func GenerateShortUrl(c *gin.Context) {
	var request UrlCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := services.CreateShortUrl(request.OriginalUrl, request.UserId)
	services.SaveUrlMapping(shortUrl, request.OriginalUrl)

	c.JSON(200, gin.H{
		"short_url": fmt.Sprintf("%s/%s", services.GetAppApiAddr(), shortUrl),
	})
}

// RedirectShortUrl reads the original url from the short ID and sends a redirect response
func RedirectShortUrl(c *gin.Context) {
	shortUrl := c.Param("short_url")
	originalUrl := services.GetUrlMapping(shortUrl)

	c.Redirect(302, originalUrl)
}
