package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/models"
	"github.com/indigowar/services/orders/storage/gen"
)

//go:generate sqlc generate .

type OrderStorage struct {
	connection *pgx.Conn
}

var _ commands.OrderStorage = &OrderStorage{}

// Add implements commands.OrderStorage.
func (o *OrderStorage) Add(ctx context.Context, order models.Order) error {
	handler := func(ctx context.Context, tx pgx.Tx) error {
		queries := gen.New(tx)
		orderId := pgtype.UUID{Bytes: order.ID, Valid: true}

		if _, err := queries.InsertOrder(ctx, gen.InsertOrderParams{
			ID:              orderId,
			Restaurant:      pgtype.UUID{Bytes: order.Restaurant, Valid: true},
			CreatedAt:       pgtype.Timestamp{Time: order.CreatedAt, Valid: true},
			Customer:        pgtype.UUID{Bytes: order.Customer.ID, Valid: true},
			CustomerAddress: order.Customer.Address,
		}); err != nil {
			return err
		}

		return nil
	}

	return runInTx(ctx, o.connection, handler)
}

// Delete implements commands.OrderStorage.
func (o *OrderStorage) Delete(ctx context.Context, id uuid.UUID) error {
	queries := gen.New(o.connection)

	if _, err := queries.DeleteOrderByID(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		return mapNoRowsError(err, "order", "id")
	}

	return nil
}

// Get implements commands.OrderStorage.
func (o *OrderStorage) Get(ctx context.Context, id uuid.UUID) (models.Order, error) {
	handler := func(ctx context.Context, tx pgx.Tx) (models.Order, error) {
		queries := gen.New(tx)
		orderId := pgtype.UUID{Bytes: id, Valid: true}

		order, err := fetchOrder(ctx, queries, orderId)
		if err != nil {
			return models.Order{}, err
		}

		optionalFieldFetchers := []func(context.Context, *gen.Queries, pgtype.UUID, *models.Order) error{
			fetchAndApplyOrderAcceptance,
			fetchAndApplyOrderCourier,
			fetchAndApplyOrderPayment,
			fetchAnyApplyOrderCancellation,
			fetchAndApplyOrderProducts,
		}

		for _, fetcher := range optionalFieldFetchers {
			err := fetcher(ctx, queries, orderId, &order)
			if err != nil {
				return models.Order{}, err
			}
		}

		return order, nil
	}

	return runInTxWithReturn(ctx, o.connection, handler)
}

// Update implements commands.OrderStorage.
func (o *OrderStorage) Update(ctx context.Context, order models.Order) error {
	handler := func(ctx context.Context, tx pgx.Tx) error {
		queries := gen.New(tx)

		if _, err := queries.UpdateOrderByID(ctx, gen.UpdateOrderByIDParams{
			ID:              pgtype.UUID{Bytes: order.ID, Valid: true},
			Restaurant:      pgtype.UUID{Bytes: order.Restaurant, Valid: true},
			CreatedAt:       pgtype.Timestamp{Time: order.CreatedAt, Valid: true},
			Customer:        pgtype.UUID{Bytes: order.Customer.ID, Valid: true},
			CustomerAddress: order.Customer.Address,
		}); err != nil {
			return mapNoRowsError(err, "order", "id")
		}

		updaters := []func(context.Context, *gen.Queries, *models.Order) error{
			insertOrUpdateOrderAcceptance,
			insertOrUpdateOrderCourier,
			insertOrUpdateOrderPayment,
			insertOrUpdateOrderCancellation,
		}

		for _, updater := range updaters {
			if err := updater(ctx, queries, &order); err != nil {
				return err
			}
		}

		return nil
	}

	return runInTx(ctx, o.connection, handler)
}

func NewOrderStorage(conn *pgx.Conn) OrderStorage {
	return OrderStorage{
		connection: conn,
	}
}
