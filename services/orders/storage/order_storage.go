package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/services/orders/commands"
	"github.com/indigowar/services/orders/models"
	"github.com/indigowar/services/orders/storage/gen"
)

//go:generate sqlc generate .

type OrderStrorage struct {
	connection *pgx.Conn
	queries    *gen.Queries
}

var _ commands.OrderStorage = &OrderStrorage{}

// Add implements commands.OrderStorage.
func (o *OrderStrorage) Add(ctx context.Context, order models.Order) error {
	panic("unimplemented")
}

// Delete implements commands.OrderStorage.
func (o *OrderStrorage) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Get implements commands.OrderStorage.
func (o *OrderStrorage) Get(ctx context.Context, id uuid.UUID) (models.Order, error) {
	panic("unimplemented")
}

// Update implements commands.OrderStorage.
func (o *OrderStrorage) Update(ctx context.Context, order models.Order) error {
	panic("unimplemented")
}

func NewOrderStorage(conn *pgx.Conn) OrderStrorage {
	return OrderStrorage{
		connection: conn,
		queries:    gen.New(conn),
	}
}
