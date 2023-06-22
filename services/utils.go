package services

import (
	"fmt"
	"os"
)

func GetAppApiAddr() string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	address := fmt.Sprintf("%s:%s", host, port)
	return address
}
