package cache

import (
	"context"
	"fmt"
	"myServer/log"

	"myServer/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var RedisClient *redis.Client
var RedisContext = context.Background()

func InitCache() {
	rConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Username: rConfig.RedisUsername,
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})
	_, err := client.Ping(RedisContext).Result()
	if err != nil {
		log.InfoLog("ping redis failed", zap.Error(err))
		panic(err)
	}
	RedisClient = client
}

func RedisTest() {
	err := RedisClient.Set(RedisContext, "jiaoname", "xiaojiao", 0).Err()
	if err != nil {
		log.InfoLog("set redis failed", zap.Error(err))
	}
	value, err := RedisClient.Get(RedisContext, "jiaoname").Result()
	fmt.Println(value)
	if err != nil {
		log.InfoLog("get redis failed", zap.Error(err))
	}
	log.InfoLog("get redis value", zap.String("value", value))
}
