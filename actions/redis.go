package actions

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
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

func Get(key string) (int, error) {
	rdb := initRedisClient()
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	newValue, _ := strconv.Atoi(value)
	return newValue, nil
}

func Set(key string, value int) {
	rdb := initRedisClient()
	rdb.Set(ctx, key, value, 0)
}

func Decrease(key string, dmg int) {
	value, _ := Get(key)
	value = value - dmg
	Set(key, value)
}

func Increase(key string, heal int) {
	value, _ := Get(key)
	value = value + heal
	Set(key, value)
}
