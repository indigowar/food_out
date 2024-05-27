package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

// mockRetrieveByID mocks RetrieveByID(context.Context, uuid.UUID) (*domain.Dish, error)
func mockRetrieveByID(dish *domain.Dish, err error) func(context.Context, uuid.UUID) (*domain.Dish, error) {
	return func(_ context.Context, _ uuid.UUID) (*domain.Dish, error) {
		return dish, err
	}
}

// mockRetrieveByMenu mocks RetrieveByMenu(context.Context, uuid.UUID) ([]*domain.Dish, error)
func mocksRetrieveByMenu(dishes []*domain.Dish, err error) func(context.Context, uuid.UUID) ([]*domain.Dish, error) {
	return func(_ context.Context, _ uuid.UUID) ([]*domain.Dish, error) {
		return dishes, err
	}
}

// mockRetrieveByRestaurant mocks RetrieveByRestaurant(context.Context, uuid.UUID) ([]*domain.Dish, error)
func mockRetrieveByRestaurant(dishes []*domain.Dish, err error) func(context.Context, uuid.UUID) ([]*domain.Dish, error) {
	return func(_ context.Context, _ uuid.UUID) ([]*domain.Dish, error) {
		return dishes, err
	}
}
