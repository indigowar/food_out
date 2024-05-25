package commands

import (
	"context"
	"net/url"

	"github.com/google/uuid"
)

type UpdateDishArgument struct {
	Id    uuid.UUID
	Name  string
	Image *url.URL
	Price float64
}

type UpdateDishCommand struct{}

func (cmd *UpdateDishCommand) UpdateDish(ctx context.Context, arg UpdateDishArgument) error {
	panic("not implemented")
}

func NewUpdateDishCommand() *UpdateDishCommand {
	panic("not implemented")
}
