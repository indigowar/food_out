package commands

import (
	"context"

	"github.com/google/uuid"
)

type DeleteMenuCommand struct{}

func (cmd *DeleteMenuCommand) DeleteMenu(ctx context.Context, restaurant uuid.UUID, menu uuid.UUID) error {
	panic("not implemented")
}

func NewDeleteMenuCommand() *DeleteMenuCommand {
	panic("not implemented")
}
