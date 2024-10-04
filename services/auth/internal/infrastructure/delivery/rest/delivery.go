package rest

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest/gen"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

type Delivery struct {
	logger *slog.Logger
	server *http.Server
}

func (d *Delivery) Run(ctx context.Context) error {
	errChan := make(chan error, 1)

	go func() {
		if err := d.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			d.logger.Warn("server.ListenAndServe failed", "err", err)
			errChan <- err
		}
	}()

	d.logger.Info("rest.Server is started")

	select {
	case err := <-errChan:
		return fmt.Errorf("rest.Delivery error: %w", err)
	case <-ctx.Done():
		cancelCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := d.server.Shutdown(cancelCtx); err != nil {
			d.logger.Warn("rest.Delivery: server shutdown failed", "err", err)
			return fmt.Errorf("rest.Delivery: Shutdown: %w", err)
		}

		return nil
	}
}

func New(
	logger *slog.Logger,
	service *service.Service,
	addr string,
	accessTokenDuration time.Duration,
	authSecret []byte,
) (Delivery, error) {
	wrapper := &Wrapper{
		sv: service,
		tk: NewJwtTokenGenerator(authSecret, accessTokenDuration),
	}

	refreshSecurityHandler := &RefreshSecurityHandler{
		sv: service,
	}

	api, err := gen.NewServer(wrapper, refreshSecurityHandler)
	if err != nil {
		return Delivery{}, fmt.Errorf("rest.Delivery failed to create api: %w", err)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: api,
	}

	return Delivery{
		logger: logger,
		server: server,
	}, nil
}
