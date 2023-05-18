package initialize

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"

	"tutu-gin/core/global"
)

func InitRedis() {
	config := global.SERVICE_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Panicln("Redis Error: ", err.Error())
	}
	global.REDIS_CLIENT = client

	client2 := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       1,
	})
	err = client.Ping(context.Background()).Err()
	if err != nil {
		log.Panicln("Redis Error: ", err.Error())
	}
	global.REDIS_CLIENT = client2
}
