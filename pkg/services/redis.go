package services

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"time"
)

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		ReadTimeout: time.Duration(10 * time.Second),
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	return client
}