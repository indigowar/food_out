package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/jackc/pgx/v5"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/config"
	"github.com/indigowar/services/orders/eventconsumers"
	"github.com/indigowar/services/orders/eventproducers"
	"github.com/indigowar/services/orders/storage"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("Service configuration is invalid", "err", err)
		os.Exit(1)
	}

	postgresCon, err := connectToPostgres(cfg.Db)
	if err != nil {
		logger.Error("Failed to connect to PostgreSQL", "err", err)
		os.Exit(1)
	}

	orderStorage := storage.NewOrderStorage(postgresCon)

	orderEndedProducer := eventproducers.NewOrderEndedProducer(
		cfg.MessageQueue.Topics.OrderEnded,
		cfg.MessageQueue.Hosts,
	)

	acceptOrder, err := eventconsumers.NewAcceptOrderConsumer(
		logger,
		commands.NewAcceptOrder(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.AcceptOrder,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "accept_order", "err", err)
		os.Exit(1)
	}

	createOrder, err := eventconsumers.NewCreateOrderConsumer(
		logger,
		commands.NewCreateOrder(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.CreateOrder,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "create_order", "err", err)
		os.Exit(1)
	}

	cookingStarted, err := eventconsumers.NewCookingStartedConsumer(
		logger,
		commands.NewMarkCookingStarted(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.CookingStarted,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "cooking_started", "err", err)
		os.Exit(1)
	}

	deliveryCompleted, err := eventconsumers.NewDeliveryCompletedConsumer(
		logger,
		commands.NewMarkDeliveryCompleted(logger, &orderStorage, orderEndedProducer),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.DeliveryCompleted,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "delivery_completed", "err", err)
		os.Exit(1)
	}

	deliveryStarted, err := eventconsumers.NewDeliveryStartedConsumer(
		logger,
		commands.NewMarkDeliveryStarted(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.DeliveryStarted,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "delivery_started", "err", err)
		os.Exit(1)
	}

	orderPayed, err := eventconsumers.NewOrderPayedConsumer(
		logger,
		commands.NewPayForOrder(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.OrderHasBeenPayed,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "order_payed", "err", err)
		os.Exit(1)
	}

	takeOrder, err := eventconsumers.NewTakeOrderConsumer(
		logger,
		commands.NewTakeOrder(logger, &orderStorage),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.OrderHasBeenPayed,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "take_order", "err", err)
		os.Exit(1)
	}

	cancelOrder, err := eventconsumers.NewCancelOrderConsumer(
		logger,
		commands.NewCancelOrder(logger, &orderStorage, orderEndedProducer),
		cfg.MessageQueue.Hosts,
		cfg.MessageQueue.Group,
		cfg.MessageQueue.Topics.CancelOrder,
		0,
	)
	if err != nil {
		logger.Error("Failed to create a consumer", "event", "cancel_order", "err", err)
		os.Exit(1)
	}

	run(
		logger,
		[]eventconsumers.Consumer{
			acceptOrder,
			createOrder,
			cookingStarted,
			deliveryCompleted,
			deliveryStarted,
			orderPayed,
			takeOrder,
			cancelOrder,
		},
	// api,
	)

}

func connectToPostgres(cfg config.Postgres) (*pgx.Conn, error) {
	postgresUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Db,
	)
	return pgx.Connect(context.Background(), postgresUrl)
}

func run(
	logger *slog.Logger,
	consumers []eventconsumers.Consumer,
	// api api.API,
) {
	runCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	var wg sync.WaitGroup
	for _, consumer := range consumers {
		wg.Add(1)
		go func() {
			err := consumer.Run(runCtx)
			if err != nil {
				logger.Error("Consumer failed", "err", err)
				cancel()
			}
			wg.Done()
		}()
	}

	// wg.Add(1)
	// go func() {
	// 	err := api.Run(runCtx)
	// 	if err != nil {
	// 		logger.Error("API failure", "err", err)
	// 		cancel()
	// 	}
	// 	wg.Done()
	// }()

	wg.Wait()
}
