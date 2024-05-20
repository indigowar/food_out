package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}

	Security struct {
		Key string
	}

	Kafka struct {
		Host                string
		Port                int
		AccountCreatedTopic string
		AccountDeletedTopic string
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	if err := loadPostgres(cfg); err != nil {
		return nil, err
	}

	if err := loadSecurity(cfg); err != nil {
		return nil, err
	}

	if err := loadKafka(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadPostgres(cfg *Config) error {
	var err error

	if cfg.Postgres.Host, err = loadString("POSTGRES_HOST"); err != nil {
		return err
	}

	if cfg.Postgres.Port, err = loadNumber("POSTGRES_PORT"); err != nil {
		return err
	}

	if cfg.Postgres.User, err = loadString("POSTGRES_USER"); err != nil {
		return err
	}

	if cfg.Postgres.Password, err = loadString("POSTGRES_PASSWORD"); err != nil {
		return err
	}

	if cfg.Postgres.Database, err = loadString("POSTGRES_DB"); err != nil {
		return err
	}

	return nil
}

func loadSecurity(cfg *Config) error {
	var err error
	if cfg.Security.Key, err = loadString("SECURITY_KEY"); err != nil {
		return err
	}
	return nil
}

func loadKafka(cfg *Config) error {
	var err error

	if cfg.Kafka.Host, err = loadString("KAFKA_HOST"); err != nil {
		return err
	}

	if cfg.Kafka.Port, err = loadNumber("KAFKA_PORT"); err != nil {
		return err
	}

	if cfg.Kafka.AccountCreatedTopic, err = loadString("KAFKA_ACCOUNT_CREATED_TOPIC"); err != nil {
		return err
	}

	if cfg.Kafka.AccountDeletedTopic, err = loadString("KAFKA_ACCOUNT_DELETED_TOPIC"); err != nil {
		return err
	}

	return nil
}

func loadString(env string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		return "", fmt.Errorf("%s is not defined", env)
	}
	return value, nil
}

func loadNumber(env string) (int, error) {
	value := os.Getenv(env)
	if value == "" {
		return 0, fmt.Errorf("%s is not defined", env)
	}
	number, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("%s is invalid", env)
	}
	return int(number), nil
}
