package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetDishesByMenuQuery struct{}

func (q *GetDishesByMenuQuery) GetDishesByMenu(ctx context.Context, id uuid.UUID) ([]domain.Dish, error) {
	panic("not implemented")
}

func NewGetDishesByMenuQuery() *GetDishesByMenuQuery {
	panic("not implemented")
}
