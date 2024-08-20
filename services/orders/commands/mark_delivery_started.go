package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type MarkDeliveryStarted struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *MarkDeliveryStarted) MarkDeliveryStarted(
	ctx context.Context,
	id uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("mark delivery started: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "MarkDeliveryStarted", "Get", err)
		return ErrUnexpected
	}

	if order.DeliveryStartedAt != nil {
		return fmt.Errorf("mark delivery started: %w", ErrActionAlreadyDone)
	}

	order.DeliveryStartedAt = &timestamp

	if err := cmd.storage.Update(ctx, order); err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "MarkDeliveryStarted", "Update", err)
		return ErrUnexpected
	}

	return nil
}

func NewMarkDeliveryStarted(
	logger *slog.Logger,
	storage OrderStorage,
) *MarkDeliveryStarted {
	return &MarkDeliveryStarted{
		logger:  logger,
		storage: storage,
	}
}
