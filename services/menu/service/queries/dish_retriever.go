package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

//go:generate go run github.com/matryer/moq -out dish_retriever_moq_test.go . DishRetriever

// DishRetriever - a port to the storage to retrieve the dishes
type DishRetriever interface {
	RetrieveByID(ctx context.Context, id uuid.UUID) (*domain.Dish, error)
	RetrieveByMenu(ctx context.Context, menu uuid.UUID) ([]*domain.Dish, error)
	RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Dish, error)
}
