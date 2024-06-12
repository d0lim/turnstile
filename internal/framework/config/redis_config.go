package config

import (
	"github.com/gofiber/storage/redis/v3"
	"os"
)

type RedisConfig struct {
	Store *redis.Storage
}

func NewRedisConfig() *RedisConfig {
	store := redis.New(redis.Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     6379,
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	return &RedisConfig{
		Store: store,
	}
}
