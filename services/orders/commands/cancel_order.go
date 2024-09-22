package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/indigowar/services/orders/models"
)

type CancelOrder struct {
	logger   *slog.Logger
	storage  OrderStorage
	producer OrderEndedProducer
}

func (cmd *CancelOrder) CancelOrder(
	ctx context.Context,
	orderId uuid.UUID,
	canceller uuid.UUID,
	timestamp time.Time,
) error {
	if timestamp.After(time.Now()) {
		return fmt.Errorf("cancel order: %w: timestamp", ErrInvalidRequest)
	}

	order, err := cmd.storage.Get(ctx, orderId)
	if err != nil {
		var storageErr *StorageError
		if errors.As(err, &storageErr) {
			if storageErr.ErrorType == StorageErrorTypeNotFound {
				return ErrOrderNotFound
			}
		}

		logStorageError(cmd.logger, "CancelOrder", "Get", err)
		return ErrUnexpected
	}

	if order.Customer.ID == canceller {
		return cmd.cancelOrder(ctx, order, canceller, timestamp)
	}

	if order.Restaurant == canceller {
		return cmd.cancelOrder(ctx, order, canceller, timestamp)
	}

	if order.Courier != nil {
		if order.Courier.ID == canceller {
			return cmd.cancelOrder(ctx, order, canceller, timestamp)
		}
	}

	return fmt.Errorf("cancel order: %w: canceller", ErrInvalidRequest)
}

func (cmd *CancelOrder) cancelOrder(ctx context.Context, order models.Order, canceller uuid.UUID, timestamp time.Time) error {
	order.Cancellation = &struct {
		Canceller   uuid.UUID
		CancelledAt time.Time
	}{
		Canceller:   canceller,
		CancelledAt: timestamp,
	}

	if err := cmd.producer.Produce(ctx, order); err != nil {
		return ErrUnexpected
	}

	if err := cmd.storage.Delete(ctx, order.ID); err != nil {
		logStorageError(cmd.logger, "CancelOrder", "Delete", err)
		return ErrUnexpected
	}

	return nil
}

func NewCancelOrder(logger *slog.Logger, storage OrderStorage, producer OrderEndedProducer) *CancelOrder {
	return &CancelOrder{
		logger:   logger,
		storage:  storage,
		producer: producer,
	}
}
