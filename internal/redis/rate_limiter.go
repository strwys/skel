package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

func (cache *redisCache) Allow() bool {
	err := cache.increment()
	if err != nil {
		log.Println("Error incrementing sliding window:", err)
		return false
	}
	err = cache.removeExpired()
	if err != nil {
		log.Println("Error removing expired sliding window entries:", err)
		return false
	}
	count, err := cache.countRequests()
	if err != nil {
		log.Println("Error counting sliding window requests:", err)
		return false
	}
	allowedRequests := int64(cache.rateLimiter * cache.slidingWindow.Seconds())
	return count <= allowedRequests
}

func (cache *redisCache) increment() error {
	now := time.Now().UnixNano()
	score := float64(now)
	member := fmt.Sprintf("%d", now)
	_, err := cache.client.ZAdd(RedisPrefixRateLimiter, redis.Z{
		Score:  score,
		Member: member,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}

func (cache *redisCache) removeExpired() error {
	now := time.Now().UnixNano()
	minScore := float64(now) - cache.slidingWindow.Seconds()*1e9
	_, err := cache.client.ZRemRangeByScore(RedisPrefixRateLimiter, "0", fmt.Sprintf("%.0f", minScore)).Result()
	if err != nil {
		return err
	}
	return nil
}

func (cache *redisCache) countRequests() (int64, error) {
	now := time.Now().UnixNano()
	minScore := float64(now) - cache.slidingWindow.Seconds()*1e9
	count, err := cache.client.ZCount(RedisPrefixRateLimiter, fmt.Sprintf("%.0f", minScore), "+inf").Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
