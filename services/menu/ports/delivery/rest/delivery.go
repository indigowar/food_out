package rest

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/indigowar/food_out/services/menu/ports/delivery/rest/api"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type Delivery struct {
	logger *slog.Logger

	server *http.Server
}

func (d *Delivery) Run(ctx context.Context) error {
	errChan := make(chan error, 1)

	go func() {
		err := d.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			d.logger.Warn("Server.ListenAndServe failed", "err", err)
			errChan <- err
		}
	}()

	d.logger.Info("Rest Server is started")

	select {
	case err := <-errChan:
		return fmt.Errorf("Delivery error: Listening: %w", err)
	case <-ctx.Done():
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := d.server.Shutdown(stopCtx); err != nil {
			d.logger.Warn("Server.Shutdown failed, forcing the shutdown", "err", err)
			return fmt.Errorf("Delivery error: Shutdown: %w", err)
		}
		return nil
	}
}

func New(
	logger *slog.Logger,
	dishById queries.GetDishByIDQuery,
	menuById queries.GetMenuByIDQuery,
	menuByRestaurant queries.GetRestaurantsMenusQuery,
	restaurantList queries.GetRestaurantsQuery,
	dishValidation queries.ValidateDishOwnershipQuery,
) (*Delivery, error) {
	api, err := api.NewServer(&handler{
		dishById:         dishById,
		menuById:         menuById,
		menuByRestaurant: menuByRestaurant,
		restaurantList:   restaurantList,
		dishValidation:   dishValidation,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create API handler: %w", err)
	}

	server := &http.Server{
		Addr:    ":80",
		Handler: api,
	}

	return &Delivery{
		logger: logger,
		server: server,
	}, nil
}
