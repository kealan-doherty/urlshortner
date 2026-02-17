package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// struct wrapper around the raw Redis Client
type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStorageService() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error Init Redis - ERROR: %v", err))
	}

	fmt.Printf("\nRedis started sucessfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// acts as the setter function to place original url and new shorten url in the Redis cache
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving url | Error: %v - shortUrl: %s, originalUrl: %s", err, shortUrl, originalUrl))
	}
}

// will get the original url from the Redis cache using the shorten url as the key
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
