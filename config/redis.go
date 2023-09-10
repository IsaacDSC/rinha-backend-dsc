package config

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var onceRedis sync.Once

func StartConnectionRedis() *redis.Client {
	env := GetEnv()
	onceRedis.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: env.REDIS_URL,
			DB:   0,
		})
		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			log.Panic("Error-Ping-Redis")
		}
	})
	return redisClient
}

func CloseConnectionRedis() {
	redisClient.Close()
}
