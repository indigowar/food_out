package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
)

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrInvalidData   = errors.New("provided data is invalid")
	ErrUnexpected    = errors.New("unexpected error occurred")
)

type OrderHistory struct {
	logger  *slog.Logger
	storage Storage
}

func (svc *OrderHistory) AddOrder(ctx context.Context, order Order) error {
	// check that products belongs to the restaurant
	for _, v := range order.Products {
		if v.Restaurant != order.Restaurant {
			return fmt.Errorf("%w: products do not belong to the restaurant", ErrInvalidData)
		}
	}

	if err := svc.storage.Add(ctx, order); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return ErrAlreadyExists
		}
		svc.logger.Warn("Add Order failed", "err", err)
		return ErrUnexpected
	}

	return nil
}

func NewOrderHistory(logger *slog.Logger, storage Storage) *OrderHistory {
	return &OrderHistory{
		logger:  logger,
		storage: storage,
	}
}
