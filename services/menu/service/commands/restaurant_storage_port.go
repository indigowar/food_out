package commands

import (
	"context"

	"github.com/google/uuid"
)

//go:generate go run github.com/matryer/moq -out restaurant_storage_port_moq_test.go . RestaurantStoragePort

// RestaurantStoragePort - is a port to restaurant storage
type RestaurantStoragePort interface {
	RestaurantExists(ctx context.Context, id uuid.UUID) (bool, error)
	AddRestaurant(ctx context.Context, id uuid.UUID) error
}
