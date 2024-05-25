package commands

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

type AddMenuArgument struct {
	Name       string
	Restaurant uuid.UUID
	Image      *url.URL
}

type AddMenuCommand struct{}

func (cmd *AddMenuCommand) AddMenu(ctx context.Context, args AddMenuArgument) (uuid.UUID, error) {
	panic("not implemented")
}

func NewAddMenuCommand() *AddMenuCommand {
	panic("not implemented")
}
