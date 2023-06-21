package routing

import "github.com/gin-gonic/gin"

func BuildRouters(r *gin.Engine) {
	r.GET("/", Index)
	r.POST("/shorten-url", GenerateShortUrl)
	r.GET("/:short_url", RedirectShortUrl)
}
