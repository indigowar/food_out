package commands

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

type UpdateMenuArgument struct {
	ID    uuid.UUID
	Name  string
	Image *url.URL
}

type UpdateMenuCommand struct{}

func (cmd *UpdateMenuCommand) UpdateMenu(ctx context.Context, arg UpdateMenuArgument) error {
	panic("not implemented")
}

func NewUpdateMenuCommand() *UpdateMenuCommand {
	panic("not implemented")
}
