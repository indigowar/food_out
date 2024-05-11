package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/indigowar/food_out/services/auth/internal/infrastructure/credentials_validator/client"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/delivery/rest/api"
	"github.com/indigowar/food_out/services/auth/internal/infrastructure/storage/redis"
	"github.com/indigowar/food_out/services/auth/internal/service"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := Load()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := redis.Connect(
		cfg.Redis.Host,
		cfg.Redis.Port,
		cfg.Redis.Username,
		cfg.Redis.Password,
		cfg.Redis.DB,
	)

	if err != nil {
		log.Fatal(err)
	}

	credentialsValidator, err := client.NewClient(cfg.Accounts.Url)
	if err != nil {
		log.Fatal(err)
	}

	sessionStorage := redis.NewStorage(redisClient)

	service := service.NewService(
		logger,
		sessionStorage,
		credentialsValidator,
		cfg.Auth.SessionDuration,
	)

	wrapper := rest.NewWrapper(
		service,
		rest.NewJwtTokenGenerator(
			cfg.Auth.Key,
			cfg.Auth.AccessTokenDuration,
		),
	)

	securityHandler := rest.NewRefreshSecurityHandler(service)

	api, err := api.NewServer(wrapper, securityHandler)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":80",
		Handler: api,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil || err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Service is stopped")
}
