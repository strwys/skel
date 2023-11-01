package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/strwys/fms/config"
)

const (
	RedisPrefixRateLimiter = `rate_limiter`
)

type RedisCache interface {
	Allow() bool
}

type redisCache struct {
	client        *redis.Client
	rateLimiter   float64
	slidingWindow time.Duration
}

// NewRedisCache return new redis cache
func NewRedisCache(client *redis.Client, cfg config.Config) (RedisCache, error) {
	return &redisCache{
		client:        client,
		rateLimiter:   cfg.Redis.RateLimiter,
		slidingWindow: time.Duration(cfg.Redis.SlidingWindow) * time.Second,
	}, nil
}

func (redis *redisCache) Close() error {
	if redis.client != nil {
		if err := redis.client.Close(); err != nil {
			return err
		}
	}
	return nil
}
