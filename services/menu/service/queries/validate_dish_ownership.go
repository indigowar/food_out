package queries

import (
	"context"

	"github.com/google/uuid"
)

type ValidateDishOwnershipQuery struct{}

func (q *ValidateDishOwnershipQuery) ValidateDishOwnership(ctx context.Context, restaurantId uuid.UUID) (bool, error) {
	panic("not implemented")
}

func NewValidateDishOwnershipQuery() *ValidateDishOwnershipQuery {
	panic("not implemented")
}
