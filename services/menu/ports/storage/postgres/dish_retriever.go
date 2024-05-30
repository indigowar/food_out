package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type DishRetriever struct {
	queries *data.Queries
}

var _ queries.DishRetriever = &DishRetriever{}

// RetrieveByID implements queries.DishRetriever.
func (d *DishRetriever) RetrieveByID(ctx context.Context, id uuid.UUID) (*domain.Dish, error) {
	result, err := d.queries.RetrieveDishByID(ctx, pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, queries.ErrNotFound
		}
		return nil, err
	}

	dish, err := dishToModel(result)
	if err != nil {
		return nil, err
	}

	return dish, nil
}

// RetrieveByMenu implements queries.DishRetriever.
func (d *DishRetriever) RetrieveByMenu(ctx context.Context, menu uuid.UUID) ([]*domain.Dish, error) {
	result, err := d.queries.RetrieveDishByMenu(ctx, pgtype.UUID{Bytes: menu, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return make([]*domain.Dish, 0), nil
		}
		return nil, err
	}

	dishes := make([]*domain.Dish, len(result), 0)
	for _, item := range result {
		dish, err := dishToModel(item)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

// RetrieveByRestaurant implements queries.DishRetriever.
func (d *DishRetriever) RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Dish, error) {
	result, err := d.queries.RetrieveDishesByRestaurant(ctx, pgtype.UUID{Bytes: restaurant, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return make([]*domain.Dish, 0), nil
		}
		return nil, err
	}

	dishes := make([]*domain.Dish, len(result), 0)
	for _, item := range result {
		dish, err := dishToModel(item)
		if err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func NewDishRetriever(conn *pgx.Conn) *DishRetriever {
	return &DishRetriever{
		queries: data.New(conn),
	}
}
