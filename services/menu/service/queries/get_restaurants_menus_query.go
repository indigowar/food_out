package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetRestaurantsMenusQuery struct{}

func (q *GetRestaurantsMenusQuery) GetRestaurantsMenus(ctx context.Context, id uuid.UUID) ([]domain.Menu, error) {
	panic("not implemented")
}

func NewGetRestaurantsMenuQuery() *GetRestaurantsMenusQuery {
	panic("not implemented")
}
