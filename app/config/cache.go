package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sync"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
	ctx         = context.Background()
)

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		addr := os.Getenv("REDIS_ADDR")
		if addr == "" {
			addr = "localhost:6379"
		}
		redisClient = redis.NewClient(&redis.Options{
			Addr: addr,
		})

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Fatalf("Could not connect to Redis: %v", err)
		}
	})
	return redisClient
}
