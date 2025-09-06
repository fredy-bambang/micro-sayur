package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(Ctx).Result()
	if err != nil {
		panic(err)
	}

	return client
}
