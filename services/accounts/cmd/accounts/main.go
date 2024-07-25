package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest"
	"github.com/indigowar/food_out/services/accounts/internal/infra/events/kafka"
	"github.com/indigowar/food_out/services/accounts/internal/infra/storage/postgres"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	postgresConnection := connectToPostgres()
	defer postgresConnection.Close(context.Background())

	accountCreatedPublisher, accountDeletedPublisher := createKafkaPublishers()

	service := service.NewService(
		logger,
		postgres.NewStorage(postgresConnection),
		accountCreatedPublisher,
		accountDeletedPublisher,
	)

	rest, err := rest.New(logger, ":80", service, []byte(os.Getenv("SECURITY_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := rest.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func connectToPostgres() *pgx.Conn {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	c, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func createKafkaPublishers() (*kafka.AccountCreatedPublisher, *kafka.AccountDeletedPublisher) {
	host := os.Getenv("KAFKA_HOST")
	port, err := strconv.Atoi(os.Getenv("KAFKA_PORT"))
	if err != nil {
		log.Fatalf("failed to parse KAFKA_PORT value: %s", err)
	}

	accountCreated := kafka.NewAccountCreatedPublisher(host, port, os.Getenv("KAFKA_ACCOUNT_CREATED_TOPIC"), 0)
	accountDeleted := kafka.NewAccountDeletedPublisher(host, port, os.Getenv("KAFKA_ACCOUNT_DELETED_TOPIC"), 0)

	return accountCreated, accountDeleted
}
