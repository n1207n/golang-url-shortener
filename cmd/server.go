package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/n1207n/golang-url-shortener/routing"
	"github.com/n1207n/golang-url-shortener/services"
)

func main() {
	r := gin.Default()

	routing.BuildRouters(r)
	services.NewStorageService()

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start up the web server - Error %v", err))
	}
}
