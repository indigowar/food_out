package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/commands"
)

type RestaurantStorage struct {
	conn    *pgx.Conn
	queries *data.Queries
}

var _ commands.RestaurantStoragePort = &RestaurantStorage{}

// AddRestaurant implements commands.RestaurantStoragePort.
func (r *RestaurantStorage) AddRestaurant(ctx context.Context, id uuid.UUID) error {
	err := r.queries.InsertRestaurant(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if isDuplicatedKeyError(err) {
			return commands.ErrAlreadyExists
		}
		return err
	}

	return nil
}

// RestaurantExists implements commands.RestaurantStoragePort.
func (r *RestaurantStorage) RestaurantExists(ctx context.Context, id uuid.UUID) (bool, error) {
	if _, err := r.queries.GetRestaurant(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewRestaurantStorage(conn *pgx.Conn) *RestaurantStorage {
	return &RestaurantStorage{
		conn:    conn,
		queries: data.New(conn),
	}
}
