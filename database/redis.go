package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis() (*redis.Client, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass,
		DB:       0,
	})

	// ping to check connection
	_, err := rdb.Ping(context.Background()).Result()

	log.Println("Connected to Redis!")

	return rdb, err
}
