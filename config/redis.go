package config

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var onceRedis sync.Once

func StartConnectionRedis() *redis.Client {
	onceRedis.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
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
