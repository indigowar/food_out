package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type MarkCookingStarted struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *MarkCookingStarted) MarkCookingStarted(
	ctx context.Context,
	id uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("mark cooking started: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "MarkCookingStarted", "Get", err)
		return ErrUnexpected
	}

	if order.CookingStartedAt != nil {
		return fmt.Errorf("mark cooking started: %w", ErrActionAlreadyDone)
	}

	order.CookingStartedAt = &timestamp

	if err := cmd.storage.Update(ctx, order); err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "MarkCookingStarted", "Update", err)
		return ErrUnexpected
	}

	return nil
}

func NewMarkCookingStarted(
	logger *slog.Logger,
	storage OrderStorage,
) *MarkCookingStarted {
	return &MarkCookingStarted{
		logger:  logger,
		storage: storage,
	}
}
