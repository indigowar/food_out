package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Redis struct {
		Host     string
		Port     int
		DB       int
		Username string
		Password string
	}
	Auth struct {
		SessionDuration     time.Duration
		AccessTokenDuration time.Duration
		Key                 []byte
	}
	Accounts struct {
		Url string
	}
}

func Load() (*Config, error) {
	cfg := &Config{}

	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	cfg.Redis.Username = os.Getenv("REDIS_USER")
	cfg.Redis.Password = os.Getenv("REDIS_PASSWORD")

	redisPort := os.Getenv("REDIS_PORT")
	port, err := strconv.Atoi(redisPort)
	if err != nil {
		return nil, fmt.Errorf("REDIS_PORT is invalid: %w", err)
	}

	redisDB := os.Getenv("REDIS_DB")
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		return nil, fmt.Errorf("REDIS_DB is invalid: %w", err)
	}

	cfg.Redis.Port = port
	cfg.Redis.DB = db

	cfg.Auth.Key = []byte(os.Getenv("AUTH_KEY"))

	sessionDuration, err := time.ParseDuration(os.Getenv("AUTH_SESSION_DURATION"))
	if err != nil {
		return nil, fmt.Errorf("AUTH_SESSION_DURATION is invalid: %w", err)
	}

	accessDuration, err := time.ParseDuration(os.Getenv("AUTH_ACCESS_TOKEN_DURATION"))
	if err != nil {
		return nil, fmt.Errorf("AUTH_ACCESS_TOKEN_DURATION is invalid: %w", err)
	}

	cfg.Auth.SessionDuration = sessionDuration
	cfg.Auth.AccessTokenDuration = accessDuration

	cfg.Accounts.Url = os.Getenv("ACCOUNTS_URL")

	return cfg, nil
}
