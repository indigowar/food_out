package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type RestaurantRetriever struct {
	queries *data.Queries
}

var _ queries.RestaurantRetriever = &RestaurantRetriever{}

// GetRestaurants implements queries.RestaurantRetriever.
func (r *RestaurantRetriever) GetRestaurants(ctx context.Context) ([]uuid.UUID, error) {
	values, err := r.queries.GetAllRestaurants(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return make([]uuid.UUID, 0), nil
		}
		return nil, err
	}

	result := make([]uuid.UUID, len(values))
	for i, v := range values {
		result[i] = v.Bytes
	}
	return result, nil
}

// RestaurantExists implements queries.RestaurantRetriever.
func (r *RestaurantRetriever) RestaurantExists(ctx context.Context, id uuid.UUID) (bool, error) {
	if _, err := r.queries.GetRestaurant(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewRestaurantRetriever(conn *pgx.Conn) *RestaurantRetriever {
	return &RestaurantRetriever{
		queries: data.New(conn),
	}
}
