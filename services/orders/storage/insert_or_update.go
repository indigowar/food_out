package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/services/orders/models"
	"github.com/indigowar/services/orders/storage/gen"
)

func insertOrUpdateOrderAcceptance(ctx context.Context, queries *gen.Queries, order *models.Order) error {
	if order.Acceptance != nil {
		if _, err := queries.InserOrUpdateOrdersAcceptance(ctx, gen.InserOrUpdateOrdersAcceptanceParams{
			OrderID:    pgtype.UUID{Bytes: order.ID, Valid: true},
			Manager:    pgtype.UUID{},
			AcceptedAt: pgtype.Timestamp{},
		}); err != nil {
			return mapNoRowsError(err, "order", "id")
		}
	}
	return nil
}

func insertOrUpdateOrderCourier(ctx context.Context, queries *gen.Queries, order *models.Order) error {
	if order.Courier != nil {
		if _, err := queries.InsertOrUpdateOrdersCourier(ctx, gen.InsertOrUpdateOrdersCourierParams{
			OrderID: pgtype.UUID{Bytes: order.ID, Valid: true},
			Courier: pgtype.UUID{Bytes: order.Courier.ID, Valid: true},
			TakenAt: pgtype.Timestamp{Time: order.Courier.TakenAt, Valid: true},
		}); err != nil {
			return mapNoRowsError(err, "order", "id")
		}
	}

	return nil
}

func insertOrUpdateOrderPayment(ctx context.Context, queries *gen.Queries, order *models.Order) error {
	if order.Payment != nil {
		if _, err := queries.InsertOrUpdateOrdersPayment(ctx, gen.InsertOrUpdateOrdersPaymentParams{
			OrderID:     pgtype.UUID{Bytes: order.ID, Valid: true},
			Transaction: order.Payment.Transaction,
			PayedAt:     pgtype.Timestamp{Time: order.Payment.PayedAt, Valid: true},
		}); err != nil {
			return mapNoRowsError(err, "order", "id")
		}
	}

	return nil
}

func insertOrUpdateOrderCancellation(ctx context.Context, queries *gen.Queries, order *models.Order) error {
	if order.Cancellation != nil {
		if _, err := queries.InsertOrUpdateOrdersCancellation(ctx, gen.InsertOrUpdateOrdersCancellationParams{
			OrderID:     pgtype.UUID{Bytes: order.ID, Valid: true},
			Canceller:   pgtype.UUID{Bytes: order.Cancellation.Canceller, Valid: true},
			CancelledAt: pgtype.Timestamp{Time: order.Cancellation.CancelledAt, Valid: true},
		}); err != nil {
			return mapNoRowsError(err, "order", "id")
		}
	}

	return nil
}
