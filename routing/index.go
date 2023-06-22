package routing

import "github.com/gin-gonic/gin"

// Index sends a test response at root url
func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}
