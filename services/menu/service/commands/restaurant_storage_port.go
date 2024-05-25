package commands

import (
	"context"

	"github.com/google/uuid"
)

//go:generate moq -out restaurant_storage_port_moq_test.go . RestaurantStoragePort

// RestaurantStoragePort - is a port to restaurant storage
type RestaurantStoragePort interface {
	AddRestaurant(ctx context.Context, id uuid.UUID) error
}
