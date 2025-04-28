package service

import (
	"async_user_service/app/config"
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
}

var (
	CacheServiceInstance *CacheService
	cacheOnce            sync.Once
)

func NewCacheService() *CacheService {
	cacheOnce.Do(func() {
		CacheServiceInstance = &CacheService{
			client: config.GetRedisClient(),
		}
	})
	return CacheServiceInstance
}

// Generic function to get data from Redis
func (s *CacheService) Get(key string, target interface{}) error {
	ctx := context.Background()
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val), target)
}

// Generic function to set data to Redis
func (s *CacheService) Set(key string, value interface{}, ttl time.Duration) error {
	ctx := context.Background()
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.client.Set(ctx, key, bytes, ttl).Err()
}
