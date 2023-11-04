package config

import "github.com/go-redis/redis"

func (cfg Config) NewRedisClient(db int, pass string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		DB:       db,
		Password: pass,
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return client, nil
}
