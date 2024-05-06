package utils

import (
	"context"
	"gin-app/common/global"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       2,
		PoolSize: 30,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		panic("Redis连接错误")
	}

	global.RDB = client
}
