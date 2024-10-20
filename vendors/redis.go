package vendors

import (
	"context"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

func InitRedis() {
	conf := config.GetConfig()

	redisClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Addr,
		Password: conf.RedisConfig.Password,
		DB:       conf.RedisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), conf.RedisConfig.Timeout)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
