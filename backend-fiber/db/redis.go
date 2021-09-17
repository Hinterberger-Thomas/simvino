package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var ctx = context.Background()

func InitRedisSalt() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       1,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	Client = client
}
