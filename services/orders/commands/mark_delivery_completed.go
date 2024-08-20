package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type MarkDeliveryCompleted struct {
	logger     *slog.Logger
	storage    OrderStorage
	orderEnded OrderEndedProducer
}

func (cmd *MarkDeliveryCompleted) MarkDeliveryCompleted(
	ctx context.Context,
	id uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("mark delivery completed: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, id)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "MarkDeliveryCompleted", "Get", err)
		return ErrUnexpected
	}

	if order.DeliveryCompleted != nil {
		return fmt.Errorf("mark delivery completed: %w", ErrActionAlreadyDone)
	}

	order.DeliveryCompleted = &timestamp

	if err := cmd.orderEnded.Produce(ctx, order); err != nil {
		return ErrUnexpected
	}

	if err := cmd.storage.Delete(ctx, order.ID); err != nil {
		logStorageError(cmd.logger, "MarkDeliveryCompleted", "Delete", err)
		return ErrUnexpected
	}

	return nil
}

func NewMarkDeliveryCompleted(
	logger *slog.Logger,
	storage OrderStorage,
	orderEnded OrderEndedProducer,
) *MarkDeliveryCompleted {
	return &MarkDeliveryCompleted{
		logger:     logger,
		storage:    storage,
		orderEnded: orderEnded,
	}
}
