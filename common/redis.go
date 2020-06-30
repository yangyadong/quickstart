package common

import (
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis() {
	var addr = "127.0.0.1:6379"
	var password = ""

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
	})
	p, err := c.Ping().Result()
	if err != nil {
		fmt.Println("redis kill")
	}
	fmt.Println(p)
	c.Do("SET", "key", "duzhenxun")
	rs := c.Do("GET", "key").Val()
	fmt.Println(rs)
	c.Close()
}
