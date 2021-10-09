package redis

import (
	"context"
	. "stencil-go/app/core/config"

	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     GConfig.Redis.Addr,
		Password: GConfig.Redis.Password,
		DB:       GConfig.Redis.DB,
		PoolSize: GConfig.Redis.PoolSize,
	})
	if GConfig.Redis.Required {
		if _, err := client.Ping(context.Background()).Result(); err != nil {
			panic(err.Error())
		}
	}

	return client
}
