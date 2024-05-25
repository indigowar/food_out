package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetDishByIDQuery struct{}

func (q *GetDishByIDQuery) GetDishByID(ctx context.Context, id uuid.UUID) (*domain.Dish, error) {
	panic("not implemented")
}

func NewGetDishByIDQuery() *GetDishByIDQuery {
	panic("not implemented")
}
