package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/indigowar/food_out/services/order_history/api"
	"github.com/indigowar/food_out/services/order_history/config"
	"github.com/indigowar/food_out/services/order_history/eventconsumers/kafka"
	"github.com/indigowar/food_out/services/order_history/service"
	"github.com/indigowar/food_out/services/order_history/storage/postgres"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	config, err := config.Load()
	if err != nil {
		logger.Error("Service configuration failure", "err", err)
		os.Exit(1)
	}

	postgresCon, err := postgres.Connect(config.Postgres)
	if err != nil {
		logger.Error("Connection failed", "target", "postgres", "err", err)
		os.Exit(1)
	}

	service := service.NewOrderHistory(logger, postgres.NewOrderStorage(postgresCon))

	eventConsumer := kafka.NewOrderEndedConsumer(service, config.Kafka)
	api := api.NewAPI()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err := eventConsumer.Consume(ctx)
		if err != nil {
			logger.Error("Consumer failed", "err", err)
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := api.Run(ctx)
		if err != nil {
			logger.Error("API failed", "err", err)
			cancel()
		}

	}()

	wg.Wait()
}
