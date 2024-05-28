package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"os"
	"time"
)

type SessionConfig struct {
	store *session.Store
}

func NewSessionStore() *SessionConfig {
	redisStore := redis.New(redis.Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     6379,
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: 0,
	})
	sessionStore := session.New(session.Config{
		Storage:    redisStore,
		Expiration: 24 * time.Hour,
	})

	return &SessionConfig{store: sessionStore}
}
