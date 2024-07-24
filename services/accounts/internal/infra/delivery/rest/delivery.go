package rest

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/indigowar/food_out/services/accounts/internal/infra/delivery/rest/api"
	"github.com/indigowar/food_out/services/accounts/internal/service"
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
			d.logger.Warn("server.ListennAndServe failed", "err", err)
			errChan <- err
		}
	}()

	d.logger.Info("rest.Server is started")

	select {
	case err := <-errChan:
		return fmt.Errorf("rest.Delivery error: Listening: %w", err)
	case <-ctx.Done():
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := d.server.Shutdown(stopCtx); err != nil {
			d.logger.Warn("rest.Delivery: server shutdown failed", "err", err)
			return fmt.Errorf("rest.Delivery: Shutdown: %w", err)
		}

		return nil
	}
}

func New(
	logger *slog.Logger,
	addr string,
	svc *service.Service,
	authKey []byte,
) (Delivery, error) {

	wrapper := &Wrapper{
		svc:    svc,
		logger: logger,
	}
	jwtSecurity := &JwtSecurityHandler{
		key: authKey,
	}

	api, err := api.NewServer(wrapper, jwtSecurity)
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
