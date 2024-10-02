package service

import (
	"context"
	"log/slog"
)

type OrderHistory struct{}

func (svc *OrderHistory) AddOrder(ctx context.Context, order Order) error {
	// TODO: Implement OrderHistory.AddOrder
	panic("unimplemented")
}

func NewOrderHistory(logger *slog.Logger, storage Storage) *OrderHistory {
	// TODO: Implement NewOrderHistory
	panic("unimplemented")
}
