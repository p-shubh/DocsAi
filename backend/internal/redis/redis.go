package redis

import (
	"context"

	"github.com/hexanecreations/docsai/backend/internal/config"
	"github.com/redis/go-redis/v9"
)

func New(cfg *config.Config) (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
