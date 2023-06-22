package routing

import "github.com/gin-gonic/gin"

// BuildRouters registers the API endpoints
func BuildRouters(r *gin.Engine) {
	r.GET("/", Index)
	r.POST("/shorten-url", GenerateShortUrl)
	r.GET("/:short_url", RedirectShortUrl)
}
