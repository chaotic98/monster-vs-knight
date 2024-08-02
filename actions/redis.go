package actions

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var (
	client *redis.Client
	once   sync.Once
)

func InitializeRedis() {
	once.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	})
}

func GetClient() *redis.Client {
	return client
}

func CloseClient() error {
	return client.Close()
}
