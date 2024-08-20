package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type AcceptOrder struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *AcceptOrder) AcceptOrder(
	ctx context.Context,
	id uuid.UUID,
	manager uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("accept order: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "AcceptOrder", "Get", err)
		return ErrUnexpected
	}

	if order.Acceptance != nil {
		return fmt.Errorf("accept order: %w", ErrActionAlreadyDone)
	}

	order.Acceptance = &struct {
		Manager    uuid.UUID
		AcceptedAt time.Time
	}{
		Manager:    manager,
		AcceptedAt: timestamp,
	}

	if err := cmd.storage.Update(ctx, order); err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "AcceptOrder", "Update", err)
		return ErrUnexpected
	}

	return nil
}

func NewAcceptOrder(
	logger *slog.Logger,
	storage OrderStorage,
) *AcceptOrder {
	return &AcceptOrder{
		logger:  logger,
		storage: storage,
	}
}
