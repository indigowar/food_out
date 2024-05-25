package commands

import (
	"context"

	"github.com/google/uuid"
)

type AddRestaurantCommand struct{}

func (cmd *AddRestaurantCommand) AddRestaurant(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	panic("not implemented")
}

func NewAddRestaurantCommand() *AddRestaurantCommand {
	panic("not implemented")
}
