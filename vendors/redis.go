package vendors

import (
	"context"
	"github.com/SamPariatIL/weather-wrapper/config"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client

func InitRedis() {
	conf := config.GetConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisConfig.Addr,
		Password: conf.RedisConfig.Password,
		DB:       conf.RedisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), conf.RedisConfig.Timeout)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis!")
}
