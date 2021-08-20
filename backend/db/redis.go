package db

import (
	"fmt"
	"simvino/config"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func initRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: config.SecretKeys.Redis_pas,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	Client = client
}
