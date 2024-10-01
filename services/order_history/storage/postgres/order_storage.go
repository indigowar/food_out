package postgres

import (
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/order_history/service"
)

// TODO: Implement OrderStorage

type OrderStorage struct {
	conn *pgx.Conn
}

var _ service.Storage = &OrderStorage{}

func NewOrderStorage(conn *pgx.Conn) *OrderStorage {
	return &OrderStorage{
		conn: conn,
	}
}
