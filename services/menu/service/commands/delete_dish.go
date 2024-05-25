package commands

import (
	"context"

	"github.com/google/uuid"
)

type DeleteDishCommand struct{}

func (cmd *DeleteDishCommand) DeleteDish(ctx context.Context, restaurant uuid.UUID, dish uuid.UUID) error {
	panic("not implemented")
}

func NewDeleteDishCommand() *DeleteDishCommand {
	panic("not implemented")
}
