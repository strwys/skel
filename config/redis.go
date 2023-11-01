package config

import "github.com/go-redis/redis"

func (cfg Config) NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		DB:       cfg.Redis.DB,
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
