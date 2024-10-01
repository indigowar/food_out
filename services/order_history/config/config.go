package config

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
	Postgres *Postgres
	Kafka    *Kafka
}

type Postgres struct {
	Host     string
	Port     uint
	User     string
	Password string
	Database string
}

type Kafka struct {
	Brokers             []string
	Group               string
	OrderEndedTopicName string
}

func Load() (*Config, error) {
	var cfg Config

	if err := loadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("env: %w", err)
	}

	return &cfg, nil
}

func loadEnv(cfg *Config) error {
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return fmt.Errorf("POSTGRES_PORT: %w", err)
	}
	cfg.Postgres.Port = uint(postgresPort)

	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Database = os.Getenv("POSTGRES_DATABASE")

	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	if len(brokers) == 0 {
		return fmt.Errorf("KAFKA_BROKERS: is empty")
	}
	for i, v := range brokers {
		format := regexp.MustCompile("^[a-zA-Z_]+:[0-9]+$")
		if !format.MatchString(v) {
			return fmt.Errorf("KAFKA_BROKERS: %d: is invalid", i)
		}
	}
	cfg.Kafka.Brokers = brokers

	cfg.Kafka.Group = os.Getenv("KAFKA_GROUP")
	cfg.Kafka.OrderEndedTopicName = os.Getenv("KAFKA_TOPIC_ORDER_ENDED")

	return nil
}
