package queries

import (
	"context"

	"github.com/google/uuid"
)

//go:generate moq -out dish_ownership_checker_moq_test.go . DishOwnershipChecker

// DishOwnershipChecker - a port to check if all dishes belongs to restaurant
type DishOwnershipChecker interface {
	Check(ctx context.Context, restaurantId uuid.UUID, dishes []uuid.UUID) (bool, error)
}
