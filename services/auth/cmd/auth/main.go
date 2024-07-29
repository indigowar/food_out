package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/indigowar/food_out/services/auth/internal/infrastructure/credentials_validator/client"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/storage/redis"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

func main() {
	secret, accessDuration, sessionDuration := parseAuthSettings()

	redisStorage := connectToRedis()
	credentialsValidator := connectToAccountService()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	service := service.NewService(logger, redisStorage, credentialsValidator, sessionDuration)

	delivery, err := rest.New(logger, service, ":80", accessDuration, secret)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := delivery.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func connectToRedis() *redis.Storage {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatalf("Invalid value for REDIS_PORT env: %s", err)
	}

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Invalid value for REDIS_DB env: %s", err)
	}

	client, err := redis.Connect(os.Getenv("REDIS_HOST"), port, os.Getenv("REDIS_PASSWORD"), db)
	if err != nil {
		log.Fatalf("Invalid value for REDIS_DB env: %s", err)
	}

	return redis.NewStorage(client)
}

func connectToAccountService() *client.CredentialsValidator {
	url := os.Getenv("ACCOUNTS_URL")

	cv, err := client.NewCredentialsValidator(url)
	if err != nil {
		log.Fatalf("Failed to connect to Accouts service %s", err)
	}

	return cv
}

func parseAuthSettings() ([]byte, time.Duration, time.Duration) {
	key := []byte(os.Getenv("AUTH_KEY"))

	accessDuration, err := time.ParseDuration(os.Getenv("AUTH_ACCESS_TOKEN_DURATION"))
	if err != nil {
		log.Fatalf("Failed to parse AUTH_ACCESS_TOKEN_DURATION: %s", err)
	}

	sessionDuration, err := time.ParseDuration(os.Getenv("AUTH_SESSION_DURATION"))
	if err != nil {
		log.Fatalf("Failed to parse AUTH_SESSION_DURATION: %s", err)
	}

	return key, accessDuration, sessionDuration
}
