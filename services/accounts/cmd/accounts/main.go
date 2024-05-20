package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest"
	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
	"github.com/indigowar/food_out/services/accounts/internal/infra/events/kafka"
	"github.com/indigowar/food_out/services/accounts/internal/infra/storage/postgres"
	"github.com/indigowar/food_out/services/accounts/internal/service"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := LoadConfig()
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	// Connection string: "postgres://username:password@localhost:5432/database_name"
	postgresUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.Host, cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	postgresConnection, err := pgx.Connect(context.Background(), postgresUrl)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	defer postgresConnection.Close(context.Background())

	postgresStorage := postgres.NewStorage(postgresConnection)

	accountCreatedPublisher := kafka.NewAccountCreatedPublisher(
		cfg.Kafka.Host,
		cfg.Kafka.Port,
		cfg.Kafka.AccountCreatedTopic,
		0,
	)
	accountDeletedPublisher := kafka.NewAccountDeletedPublisher(
		cfg.Kafka.Host,
		cfg.Kafka.Port,
		cfg.Kafka.AccountDeletedTopic,
		0,
	)

	service := service.NewService(
		logger,
		postgresStorage,
		accountCreatedPublisher,
		accountDeletedPublisher,
	)

	serviceWrapper := rest.NewWrapper(service, logger)
	securityHandler := rest.NewJwtSecurityHandler([]byte(cfg.Security.Key))

	apiServer, err := api.NewServer(serviceWrapper, securityHandler)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	server := &http.Server{
		Addr:    ":80",
		Handler: apiServer,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}
}
