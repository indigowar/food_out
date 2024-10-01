package postgres

import (
	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/order_history/config"
)

// TODO: Implement Connect function.

func Connect(
	config *config.Postgres,
) (*pgx.Conn, error) {
	panic("unimplemented")
}
