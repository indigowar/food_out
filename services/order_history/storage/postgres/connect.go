package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/indigowar/food_out/services/order_history/config"
)

func Connect(
	config *config.Postgres,
) (*pgx.Conn, error) {
	// "postgres://username:password@localhost:5432/database_name"
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		config.User, config.Password,
		config.Host, config.Port,
		config.Database,
	)

	return pgx.Connect(context.Background(), databaseUrl)
}
