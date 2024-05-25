package commands

import (
	"context"
	"log/slog"
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
type AddDishCommand struct {
	dishes      DishStoragePort
	menus       MenuStoragePort
	restaurants RestaurantStoragePort

	logger *slog.Logger
}

func (cmd *AddDishCommand) AddDish(ctx context.Context, args AddDishArgument) (uuid.UUID, error) {
	panic("not implemented")
}

func NewAddDishCommand(
	dishes DishStoragePort,
	menus MenuStoragePort,
	restaurant RestaurantStoragePort,
	logger *slog.Logger,
) *AddDishCommand {
	return &AddDishCommand{
		dishes:      dishes,
		menus:       menus,
		restaurants: restaurant,
		logger:      logger,
	}
}
