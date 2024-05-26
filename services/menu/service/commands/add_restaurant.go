package commands

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type AddRestaurantCommand struct {
	restaurants RestaurantStoragePort

	logger *slog.Logger
}

func (cmd *AddRestaurantCommand) AddRestaurant(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	empty := uuid.UUID{}
	if id == empty {
		return uuid.UUID{}, ErrInvalidData
	}

	if err := cmd.restaurants.AddRestaurant(ctx, id); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return uuid.UUID{}, ErrRestaurantAlreadyExists
		}
		cmd.logger.Warn(
			"AddRestaurant Command FAILED",
			"action", "AddRestaurant",
			"reason", "RestaurantStoragePort",
			"error", err.Error(),
		)

		return uuid.UUID{}, ErrInternal
	}

	return id, nil
}

func NewAddRestaurantCommand(
	restaurants RestaurantStoragePort,
	logger *slog.Logger,
) *AddRestaurantCommand {
	return &AddRestaurantCommand{
		restaurants: restaurants,
		logger:      logger,
	}
}
