package services

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

// Compilation check
var _ StorageService = StorageService{}

var (
	storageService *StorageService
	ctx            = context.Background()
)

func NewStorageService(redisAddr string, redisPassword string, redisDB int) *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initializing redis = {%s}", pong))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

	storageService = &StorageService{
		redisClient: redisClient,
	}

	return storageService
}

// SaveUrlMapping persists the url shortening record to the redis
func SaveUrlMapping(shortUrl string, originalUrl string) {
	err := storageService.redisClient.Set(ctx, shortUrl, originalUrl, -1).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error %v - shortUrl: %s - originalUrl: %s", err, shortUrl, originalUrl))
	}
}

// GetUrlMapping loads up the original url to redirect from short url
func GetUrlMapping(shortUrl string) string {
	result, err := storageService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed getting key url | Error %v - shortUrl: %s", err, shortUrl))
	}

	return result
}
