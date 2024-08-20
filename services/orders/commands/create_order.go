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

type CreateOrder struct {
	logger  *slog.Logger
	storage OrderStorage
}

func (cmd *CreateOrder) CreateOrder(
	ctx context.Context,
	customer uuid.UUID,
	customerAddress string,
	restaurant uuid.UUID,
	products []models.Product,
	timestamp time.Time,
) error {
	if !cmd.validateProducts(restaurant, products) {
		return fmt.Errorf("create order: %w: products", ErrInvalidRequest)
	}

	if !time.Now().After(timestamp) {
		return fmt.Errorf("create order: %w: createdAt", ErrInvalidRequest)
	}

	// create an order
	order := models.Order{
		ID:         uuid.New(),
		Restaurant: restaurant,
		Products:   products,
		Customer: struct {
			ID      uuid.UUID
			Address string
		}{
			ID:      customer,
			Address: customerAddress,
		},
		CreatedAt: timestamp,
	}

	if err := cmd.storage.Add(ctx, order); err != nil {
		var storageError *StorageError
		if errors.As(err, &storageError) {
			switch storageError.ErrorType {
			case StorageErrorTypeAlreadyExists:
				return ErrOrderDuplicated
			}
		}

		logStorageError(cmd.logger, "CreateOrder", "Add", err)
		return ErrUnexpected
	}

	return nil
}

func (cmd *CreateOrder) validateProducts(restaurant uuid.UUID, products []models.Product) bool {
	for _, v := range products {
		if v.Restaurant != restaurant {
			return false
		}
	}
	return true
}

func NewCreateOrder(logger *slog.Logger, storage OrderStorage) *CreateOrder {
	return &CreateOrder{
		logger:  logger,
		storage: storage,
	}
}
