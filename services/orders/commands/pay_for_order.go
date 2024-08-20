package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type PayForOrder struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *PayForOrder) PayForOrder(
	ctx context.Context,
	id uuid.UUID,
	transaction string,
	timestamp time.Time,
) error {
	if !timestamp.Before(time.Now()) {
		return fmt.Errorf("pay for order: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "PayForOrder", "Get", err)
		return ErrUnexpected
	}

	if order.Payment != nil {
		return ErrActionAlreadyDone
	}

	order.Payment = &struct {
		Transaction string
		PayedAt     time.Time
	}{
		Transaction: transaction,
		PayedAt:     timestamp,
	}

	if err := cmd.storage.Update(ctx, order); err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}

			if storageErr.ErrorType == StorageErrorTypeAlreadyExists {
				cmd.logger.Error(
					"Duplicated in the storage",
					"Command", "PayForOrder",
					"Action", "Update",
					"Object", storageErr.Object,
					"Field", storageErr.Field,
					"Err", err,
				)

				return ErrInvalidRequest
			}
		}

		logStorageError(cmd.logger, "PayForOrder", "Update", err)
		return ErrUnexpected
	}

	return nil
}

func NewPayForOrder(
	logger *slog.Logger,
	storage OrderStorage,
) *PayForOrder {
	return &PayForOrder{
		logger:  logger,
		storage: storage,
	}
}
