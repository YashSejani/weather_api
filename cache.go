package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

var rdb *redis.Client
var ctx = context.Background()

func initRedis() {
	redisURL := os.Getenv("REDIS_URL")
    if redisURL == "" {
        redisURL = "redis://localhost:6379"
    }

	options, err := redis.ParseURL(redisURL)
    if err != nil {
        panic(err)
    }

    rdb = redis.NewClient(options)

	// rdb = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })
}