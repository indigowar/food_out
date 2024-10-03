package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/order_history/service"
	"github.com/indigowar/food_out/services/order_history/storage/postgres/gen"
)

//go:generate sqlc generate .

// TODO: Implement OrderStorage

type OrderStorage struct {
	conn *pgx.Conn
}

var _ service.Storage = &OrderStorage{}

// Add implements service.Storage.
func (storage *OrderStorage) Add(ctx context.Context, order service.Order) error {
	runner := func(ctx context.Context, tx pgx.Tx) error {
		queries := gen.New(tx)

		if err := queries.InsertOrder(ctx, gen.InsertOrderParams{
			ID:                  pgtype.UUID{Bytes: order.ID, Valid: true},
			Restaurant:          pgtype.UUID{Bytes: order.Restaurant, Valid: true},
			CreatedAt:           pgtype.Timestamp{Time: order.CreatedAt, Valid: true},
			Customer:            pgtype.UUID{Bytes: order.Customer.ID, Valid: true},
			CustomerAddress:     order.Customer.Address,
			CookingStartedAt:    optionalTimeToTimestamp(order.CookingStartedAt),
			DeliveryStartedAt:   optionalTimeToTimestamp(order.DeliveryStartedAt),
			DeliveryCompletedAt: optionalTimeToTimestamp(order.DeliveryCompleted),
		}); err != nil {
			return err
		}

		if ac := order.Acceptance; ac != nil {
			if err := queries.InsertOrdersAcceptance(ctx, gen.InsertOrdersAcceptanceParams{
				OrderID:    pgtype.UUID{Bytes: order.ID, Valid: true},
				Manager:    pgtype.UUID{Bytes: ac.Manager, Valid: true},
				AcceptedAt: pgtype.Timestamp{Time: ac.AcceptedAt, Valid: true},
			}); err != nil {
				return err
			}
		}

		if c := order.Courier; c != nil {
			if err := queries.InsertOrdersCourier(ctx, gen.InsertOrdersCourierParams{
				OrderID: pgtype.UUID{Bytes: order.ID, Valid: true},
				Courier: pgtype.UUID{Bytes: c.ID, Valid: true},
				TakenAt: pgtype.Timestamp{Time: c.TakenAt, Valid: true},
			}); err != nil {
				return err
			}
		}

		if p := order.Payment; p != nil {
			if err := queries.InsertOrdersPayment(ctx, gen.InsertOrdersPaymentParams{
				OrderID:     pgtype.UUID{Bytes: order.ID, Valid: true},
				Transaction: p.Transaction,
				PayedAt:     pgtype.Timestamp{Time: p.PayedAt, Valid: true},
			}); err != nil {
				return err
			}
		}

		if c := order.Cancellation; c != nil {
			if err := queries.InsertOrdersCancellation(ctx, gen.InsertOrdersCancellationParams{
				OrderID:     pgtype.UUID{Bytes: order.ID, Valid: true},
				Canceller:   pgtype.UUID{Bytes: c.Canceller, Valid: true},
				CancelledAt: pgtype.Timestamp{Time: c.CancelledAt, Valid: true},
			}); err != nil {
				return err
			}
		}

		return nil
	}

	return executeInTransaction(ctx, storage.conn, runner)
}

func NewOrderStorage(conn *pgx.Conn) *OrderStorage {
	return &OrderStorage{
		conn: conn,
	}
}
