package commands

import (
	"context"
	"log/slog"
	"net/url"

	"github.com/google/uuid"
)

type AddMenuArgument struct {
	Name       string
	Restaurant uuid.UUID
	Image      *url.URL
}

type AddMenuCommand struct {
	restaurants RestaurantStoragePort
	menus       MenuStoragePort

	logger *slog.Logger
}

func (cmd *AddMenuCommand) AddMenu(ctx context.Context, args AddMenuArgument) (uuid.UUID, error) {
	panic("not implemented")
}

func NewAddMenuCommand(
	restaurants RestaurantStoragePort,
	menus MenuStoragePort,
	logger *slog.Logger,
) *AddMenuCommand {
	return &AddMenuCommand{
		restaurants: restaurants,
		menus:       menus,
		logger:      logger,
	}
}
