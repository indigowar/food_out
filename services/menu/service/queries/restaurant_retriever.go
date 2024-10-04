package queries

import (
	"context"

	"github.com/google/uuid"
)

//go:generate go run github.com/matryer/moq -out restaurant_retriever_moq_test.go . RestaurantRetriever

// RestaurantRetriever- a port to retrieve restaurants in the service
type RestaurantRetriever interface {
	GetRestaurants(ctx context.Context) ([]uuid.UUID, error)
	RestaurantExists(ctx context.Context, id uuid.UUID) (bool, error)
}
