package queries

import (
	"context"

	"github.com/google/uuid"
)

// GetRestaurants(ctx context.Context) ([]uuid.UUID, error)
func mockGetRestaurants(ids []uuid.UUID, err error) func(context.Context) ([]uuid.UUID, error) {
	return func(ctx context.Context) ([]uuid.UUID, error) {
		return ids, err
	}
}
