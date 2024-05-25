package commands

import (
	"context"

	"github.com/google/uuid"
)

func mockRestaurantExists(v bool, e error) func(context.Context, uuid.UUID) (bool, error) {
	return func(_ context.Context, _ uuid.UUID) (bool, error) {
		return v, e
	}
}

func mockAddRestaurant(e error) func(context.Context, uuid.UUID) error {
	return func(_ context.Context, _ uuid.UUID) error {
		return e
	}
}
