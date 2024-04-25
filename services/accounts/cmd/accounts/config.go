package main

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Postgres struct {
		Host     string
		Port     uint
		User     string
		Password string
		Database string
	}

	Security struct {
		Key string
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	portStr := os.Getenv("POSTGRES_PORT")

	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		return nil, errors.New("invalid Postgres port")
	}

	cfg.Postgres.Port = uint(port)
	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Database = os.Getenv("POSTGRES_DB")

	cfg.Security.Key = os.Getenv("SECURITY_KEY")

	if cfg.Postgres.Host == "" {
		return nil, errors.New("Postgres host is empty")
	}

	if cfg.Postgres.Port == 0 {
		return nil, errors.New("Postgres port is invalid")
	}

	if cfg.Security.Key == "" {
		return nil, errors.New("Security key is empty")
	}

	return &cfg, nil
}
