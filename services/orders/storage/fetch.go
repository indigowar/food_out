package storage

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/models"
	"github.com/indigowar/services/orders/storage/gen"
)

func fetchOrder(ctx context.Context, queries *gen.Queries, orderId pgtype.UUID) (models.Order, error) {
	orderData, err := queries.SelectOrderByID(ctx, orderId)
	if err != nil {
		return models.Order{}, mapNoRowsError(err, "order", "id")
	}

	order := models.Order{
		ID:         orderData.ID.Bytes,
		Restaurant: orderData.Restaurant.Bytes,
		Customer: struct {
			ID      uuid.UUID
			Address string
		}{
			ID:      orderData.Customer.Bytes,
			Address: orderData.CustomerAddress,
		},
		CreatedAt: orderData.CreatedAt.Time,
	}

	if orderData.CookingStartedAt.Valid {
		order.CookingStartedAt = &orderData.CookingStartedAt.Time
	}
	if orderData.DeliveryStartedAt.Valid {
		order.DeliveryStartedAt = &orderData.DeliveryStartedAt.Time
	}
	if orderData.DeliveryCompletedAt.Valid {
		order.DeliveryCompleted = &orderData.DeliveryCompletedAt.Time
	}

	return order, nil
}

func fetchAndApplyOrderAcceptance(
	ctx context.Context,
	queries *gen.Queries,
	orderId pgtype.UUID,
	order *models.Order,
) error {
	data, err := queries.SelectOrderAcceptanceByID(ctx, orderId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err == nil {
		order.Acceptance = &struct {
			Manager    uuid.UUID
			AcceptedAt time.Time
		}{
			Manager:    data.Manager.Bytes,
			AcceptedAt: data.AcceptedAt.Time,
		}
	}

	return nil
}

func fetchAndApplyOrderCourier(
	ctx context.Context,
	queries *gen.Queries,
	orderId pgtype.UUID,
	order *models.Order,
) error {
	data, err := queries.SelectOrdersCourierByOrderID(ctx, orderId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return &commands.StorageError{
			ErrorType: commands.StorageErrorTypeNotFound,
			Object:    "order",
			Field:     "id",
			Message:   errors.New("order was not found by id"),
		}
	}

	if err == nil {
		order.Courier = &struct {
			ID      uuid.UUID
			TakenAt time.Time
		}{
			ID:      data.Courier.Bytes,
			TakenAt: data.TakenAt.Time,
		}
	}

	return nil
}

func fetchAndApplyOrderPayment(
	ctx context.Context,
	queries *gen.Queries,
	orderId pgtype.UUID,
	order *models.Order,
) error {
	data, err := queries.SelectOrdersPaymentByOrderID(ctx, orderId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err == nil {
		order.Payment = &struct {
			Transaction string
			PayedAt     time.Time
		}{
			Transaction: data.Transaction,
			PayedAt:     data.PayedAt.Time,
		}
	}

	return nil
}

func fetchAnyApplyOrderCancellation(
	ctx context.Context,
	queries *gen.Queries,
	orderId pgtype.UUID,
	order *models.Order,
) error {
	data, err := queries.SelectOrdersCancellationByOrderID(ctx, orderId)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err == nil {
		order.Cancellation = &struct {
			Canceller   uuid.UUID
			CancelledAt time.Time
		}{
			Canceller:   data.Canceller.Bytes,
			CancelledAt: data.CancelledAt.Time,
		}
	}

	return nil
}

func fetchAndApplyOrderProducts(
	ctx context.Context,
	queries *gen.Queries,
	orderId pgtype.UUID,
	order *models.Order,
) error {
	products, err := queries.SelectProductsByOrderID(ctx, orderId)
	if err != nil {
		return mapNoRowsError(err, "product", "order_id")
	}

	order.Products = make([]models.Product, 0, len(products))

	for _, v := range products {
		order.Products = append(order.Products, models.Product{
			ID:          v.ID.Bytes,
			Restaurant:  v.Restaurant.Bytes,
			Name:        v.Name,
			Picture:     v.Picture,
			Price:       v.Price,
			Description: v.Description,
		})
	}

	return nil
}
