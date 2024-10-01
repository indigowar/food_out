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
	originalProducts []models.OriginalProduct,
	timestamp time.Time,
) error {
	if !cmd.validateProducts(restaurant, originalProducts) {
		return fmt.Errorf("create order: %w: products", ErrInvalidRequest)
	}

	if !time.Now().After(timestamp) {
		return fmt.Errorf("create order: %w: createdAt", ErrInvalidRequest)
	}

	products := make([]models.Product, 0, len(originalProducts))
	for _, op := range originalProducts {
		products = append(products, models.Product{
			ID:          uuid.New(),
			Original:    op.ID,
			Restaurant:  op.Restaurant,
			Name:        op.Name,
			Picture:     op.Picture,
			Price:       op.Price,
			Description: op.Description,
		})
	}

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

func (cmd *CreateOrder) validateProducts(restaurant uuid.UUID, products []models.OriginalProduct) bool {
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
