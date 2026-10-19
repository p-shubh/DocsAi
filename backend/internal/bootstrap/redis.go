package bootstrap

import (
	"github.com/hexanecreations/docsai/backend/internal/redis"
)

func (a *Application) loadRedis() error {

	client, err := redis.New(a.Config)
	if err != nil {
		return err
	}

	a.Redis = client

	a.Logger.Info("Redis connected")

	return nil
}
