package actions

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
)

var RedisClient *redis.Client

func init() {
	ctx := context.Background()
	RedisClient = initRedisClient(ctx)
}

func initRedisClient(ctx context.Context) *redis.Client {
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
	ctx := context.Background()
	RedisClient.FlushAll(ctx)
}

func Get(key string) (int, error) {
	value, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	newValue, _ := strconv.Atoi(value)
	return newValue, nil
}

func Set(key string, value int) {
	RedisClient.Set(ctx, key, value, 0)
}

func Decrease(key string, dmg int) {
	value, _ := Get(key)
	value -= dmg
	Set(key, value)
}

func Increase(key string, heal int) {
	value, _ := Get(key)
	value += heal
	Set(key, value)
}
