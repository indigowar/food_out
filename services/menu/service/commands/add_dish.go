package commands

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

type AddDishArgument struct {
	Restaurant uuid.UUID
	Menu       uuid.UUID
	Name       string
	Image      *url.URL
	Price      float64
}

// AddDishCommand - a command that handles the addition of a new dish to the menu
type AddDishCommand struct{}

func (cmd *AddDishCommand) AddDish(ctx context.Context, args AddDishArgument) (uuid.UUID, error) {
	panic("not implemented")
}

func NewAddDishCommand() *AddDishCommand {
	panic("not implemented")
}
