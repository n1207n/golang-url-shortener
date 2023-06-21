package routing

import "github.com/gin-gonic/gin"

func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}
