package cache

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/echewisi/numeris_assessment/pkg/config"
)

// CacheService handles Redis operations
type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

// NewCacheService initializes a Redis client
func NewCacheService(cfg *config.Config) *CacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password, // no password set
		DB:       0,                  // use default DB
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully")
	return &CacheService{
		client: client,
		ctx:    ctx,
	}
}

// Set sets a key-value pair in the cache
func (s *CacheService) Set(key string, value interface{}, expiration time.Duration) error {
	err := s.client.Set(s.ctx, key, value, expiration).Err()
	if err != nil {
		log.Printf("Failed to set key %s in cache: %v", key, err)
		return err
	}
	log.Printf("Key %s set in cache successfully", key)
	return nil
}

// Get retrieves a value by key from the cache
func (s *CacheService) Get(key string) (string, error) {
	val, err := s.client.Get(s.ctx, key).Result()
	if err == redis.Nil {
		log.Printf("Key %s does not exist in cache", key)
		return "", nil
	} else if err != nil {
		log.Printf("Failed to get key %s from cache: %v", key, err)
		return "", err
	}
	return val, nil
}
