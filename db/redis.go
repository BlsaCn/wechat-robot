package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
	"wechat-robot/config"
)

var Redis *redis.Client
var ctx = context.Background()

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       config.RedisDB,
	})
}

func CacheGet(key string) (string, error) {
	val, err := Redis.Get(ctx, key).Result()
	// redis.Nil 或 错误，都返回err
	if err != nil {
		return "", err
	}
	return val, nil
}

func CacheSave(key, value string) {
	_ = Redis.Set(ctx, key, value, time.Hour).Err()
}
