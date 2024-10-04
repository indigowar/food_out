package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

//go:generate go run github.com/matryer/moq -out dish_storage_port_moq_test.go . DishStoragePort

// DishStoragePort - is a port to dish storage
type DishStoragePort interface {
	GetDish(ctx context.Context, id uuid.UUID) (*domain.Dish, error)
	AddDish(ctx context.Context, dish *domain.Dish) error
	DeleteDish(ctx context.Context, id uuid.UUID) error
	UpdateDish(ctx context.Context, dish *domain.Dish) error
}
