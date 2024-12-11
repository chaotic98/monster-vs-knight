package models

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisSingleton struct {
	client *redis.Client
}

var instance *RedisSingleton

func GetInstance() *RedisSingleton {
	if instance == nil {
		instance = &RedisSingleton{}
	}
	return instance
}

var ctx = context.Background()

func (rs *RedisSingleton) ConnectToRedis(host string, port int) error {
	rs.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})

	_, err := rs.client.Ping(ctx).Result()
	if err != nil {
		return err
	}

	return nil
}

//func (rs *RedisSingleton) Set(key string, value interface{}) error {
//	return rs.client.Set(key, value, 0).Err()
//}
//
//func (rs *RedisSingleton) Get(key string) (interface{}, error) {
//	return rs.client.Get(key).Result()
//}
