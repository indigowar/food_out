package queries

import (
	"context"

	"github.com/google/uuid"
)

type GetRestaurantsQuery struct{}

func (q *GetRestaurantsQuery) GetRestaurants(ctx context.Context) ([]uuid.UUID, error) {
	panic("not implemented")
}

func NewGetRestaurantsQuery() *GetRestaurantsQuery {
	panic("not implemented")
}
