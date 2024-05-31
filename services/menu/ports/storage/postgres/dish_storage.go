package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/commands"
)

type DishStorage struct {
	queries *data.Queries
}

var _ commands.DishStoragePort = &DishStorage{}

// AddDish implements commands.DishStoragePort.
func (storage *DishStorage) AddDish(ctx context.Context, dish *domain.Dish) error {
	if err := storage.queries.InsertDish(ctx, data.InsertDishParams{
		ID:    pgtype.UUID{Bytes: dish.ID(), Valid: true},
		Name:  dish.Name(),
		Image: dish.Image().String(),
		Price: dish.Price(),
	}); err != nil {
		if isDuplicatedKeyError(err) {
			return commands.ErrAlreadyExists
		}
		return err
	}

	return nil
}

// DeleteDish implements commands.DishStoragePort.
func (storage *DishStorage) DeleteDish(ctx context.Context, id uuid.UUID) error {
	if err := storage.queries.DeleteDish(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
		// todo: change error handling (maybe the sql too), because it does not return anything about affected rows
		return err
	}
	return nil
}

// GetDish implements commands.DishStoragePort.
func (storage *DishStorage) GetDish(ctx context.Context, id uuid.UUID) (*domain.Dish, error) {
	dish, err := storage.queries.RetrieveDishByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, commands.ErrDishNotFound
		}
		return nil, err
	}
	return dishToModel(dish)
}

// UpdateDish implements commands.DishStoragePort.
func (storage *DishStorage) UpdateDish(ctx context.Context, dish *domain.Dish) error {
	if err := storage.queries.UpdateDish(ctx, data.UpdateDishParams{
		ID:    pgtype.UUID{Bytes: dish.ID(), Valid: true},
		Name:  dish.Name(),
		Image: dish.Image().String(),
		Price: dish.Price(),
	}); err != nil {
		// todo: change error handling (maybe the sql too), because it does not return anything about affected rows
		return err
	}
	return nil
}

func NewDishStorage(conn *pgx.Conn) *DishStorage {
	return &DishStorage{
		queries: data.New(conn),
	}
}
