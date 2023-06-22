package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/n1207n/golang-url-shortener/routing"
	"github.com/n1207n/golang-url-shortener/services"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	address := services.GetAppApiAddr()

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASS")
	redisAddress := fmt.Sprintf("%s:%s", redisHost, redisPort)

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse REDIS_DB value - Error %v", err))
	}

	r := gin.Default()

	routing.BuildRouters(r)
	services.NewStorageService(redisAddress, redisPass, redisDb)

	err = r.Run(address)
	if err != nil {
		panic(fmt.Sprintf("Failed to start up the web server - Error %v", err))
	}
}
