package commands

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type DeleteDishCommand struct {
	dishes      DishStoragePort
	restaurants RestaurantStoragePort
	logger      *slog.Logger
}

func (cmd *DeleteDishCommand) DeleteDish(ctx context.Context, restaurant uuid.UUID, dish uuid.UUID) error {
	panic("not implemented")
}

func NewDeleteDishCommand(
	dishes DishStoragePort,
	restaurants RestaurantStoragePort,
	logger *slog.Logger,
) *DeleteDishCommand {
	return &DeleteDishCommand{
		dishes:      dishes,
		restaurants: restaurants,
		logger:      logger,
	}
}
