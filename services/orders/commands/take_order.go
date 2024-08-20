package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type TakeOrder struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *TakeOrder) TakeOrder(
	ctx context.Context,
	id uuid.UUID,
	courier uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("take order: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		cmd.logger.Warn(
			"Unexpected Storage Error",
			"Command", "TakeOrder",
			"Action", "Get",
			"Err", err,
		)
		return ErrUnexpected
	}

	if order.Courier != nil {
		return fmt.Errorf("take order: %w", ErrActionAlreadyDone)
	}

	order.Courier = &struct {
		ID      uuid.UUID
		TakenAt time.Time
	}{
		ID:      courier,
		TakenAt: timestamp,
	}

	if err := cmd.storage.Update(ctx, order); err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}
		cmd.logger.Warn(
			"Unexpected Storage Error",
			"Command", "TakeOrder",
			"Action", "Update",
			"Err", err,
		)
		return ErrUnexpected
	}

	return nil
}

func NewTakeOrder(
	logger *slog.Logger,
	storage OrderStorage,
) *TakeOrder {
	return &TakeOrder{
		logger:  logger,
		storage: storage,
	}
}
