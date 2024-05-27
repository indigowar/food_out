package queries

import (
	"context"

	"github.com/google/uuid"
)

func mockCheck(v bool, e error) func(context.Context, uuid.UUID, []uuid.UUID) (bool, error) {
	return func(_ context.Context, _ uuid.UUID, _ []uuid.UUID) (bool, error) {
		return v, e
	}
}
