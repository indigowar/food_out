package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type API struct {
	server *http.Server
}

func (api *API) Run(ctx context.Context) error {
	go func() {
		if err := api.server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.server.Shutdown(shutdownCtx); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to shut down the server: %w", err)
	}

	return nil
}

func NewAPI() *API {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	return &API{
		server: server,
	}
}
