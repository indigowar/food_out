package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

func mockGetDish(dish *domain.Dish, err error) func(context.Context, uuid.UUID) (*domain.Dish, error) {
	return func(_ context.Context, _ uuid.UUID) (*domain.Dish, error) {
		return dish, err
	}
}

func mockAddDish(err error) func(context.Context, *domain.Dish) error {
	return func(_ context.Context, _ *domain.Dish) error {
		return err
	}
}

func mockDeleteDish(err error) func(context.Context, uuid.UUID) error {
	return func(_ context.Context, _ uuid.UUID) error {
		return err
	}
}

func mockUpdateDish(err error) func(context.Context, *domain.Dish) error {
	return func(_ context.Context, _ *domain.Dish) error {
		return err
	}
}
