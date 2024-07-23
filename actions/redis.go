package actions

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func initRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

func Reset() {
	rdb := initRedisClient()
	ctx := context.Background()
	rdb.FlushAll(ctx)
}
