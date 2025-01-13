package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func StartRedis() {
	Addr := os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")
	Password := os.Getenv("REDIS_PASSWORD")
	log.Println(Addr)
	client = redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       0,
	})
	if status := client.Ping(context.Background()); status.Err() != nil {
		log.Fatalf("не получилось установить соединение с Redis: %v", status.Err())
	}

}

var client *redis.Client

func GetRedisClient() *RedisClient {
	return &RedisClient{
		Client: client,
	}
}
