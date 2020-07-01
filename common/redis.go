package common

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     ConfigInfo.Redis.Addr,
		Password: ConfigInfo.Redis.Pwd,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

}
