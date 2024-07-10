package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/menu/ports/delivery/rest"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	con, err := connectToPostgres()
	if err != nil {
		logger.Error("Failed to connect to postgres", "err", err)
		os.Exit(1)
	}

	dishRetriever := postgres.NewDishRetriever(con)
	menuRetrievier := postgres.NewMenuRetriever(con)
	restaurantRetriever := postgres.NewRestaurantRetriever(con)

	restDelivery, err := rest.New(
		logger,
		queries.NewGetDishByIDQuery(dishRetriever, logger),
		queries.NewGetMenuByIdQuery(logger, menuRetrievier),
		queries.NewGetRestaurantsMenuQuery(menuRetrievier, restaurantRetriever, logger),
		queries.NewGetRestaurantsQuery(restaurantRetriever, logger),
		queries.NewValidateDishOwnershipQuery(menuRetrievier, restaurantRetriever, logger),
	)
	if err != nil {
		logger.Error("Failed to create a REST delivery", "err", err)
		os.Exit(1)
	}

	run(logger, restDelivery)
}

func connectToPostgres() (*pgx.Conn, error) {
	// Connection string: "postgres://username:password@localhost:5432/database_name"
	postgresUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	return pgx.Connect(context.Background(), postgresUrl)
}

func run(logger *slog.Logger, rest *rest.Delivery) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := rest.Run(ctx)
		if err != nil {
			// log the error
			cancel()
		}
		wg.Done()
	}()

	wg.Wait()
}
