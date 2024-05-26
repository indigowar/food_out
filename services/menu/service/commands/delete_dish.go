package commands

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type DeleteDishCommand struct {
	dishes      DishStoragePort
	restaurants RestaurantStoragePort
	menus       MenuStoragePort
	logger      *slog.Logger
}

func (cmd *DeleteDishCommand) DeleteDish(ctx context.Context, restaurant uuid.UUID, dish uuid.UUID) error {
	empty := uuid.UUID{}
	if restaurant == empty || dish == empty {
		return ErrInvalidData
	}

	if found, err := cmd.restaurants.RestaurantExists(ctx, restaurant); !found || err != nil {
		if !found {
			return ErrRestaurantNotFound
		}
		cmd.logInternalError("RestaurantExists", "RestaurantStoragePort", err)
		return ErrInternal
	}

	if belongs, err := cmd.belongsToRestaurant(ctx, restaurant, dish); !belongs || err != nil {
		if err != nil {
			return err
		}
		return ErrDishNotFound
	}

	if err := cmd.dishes.DeleteDish(ctx, dish); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrDishNotFound
		}
		cmd.logInternalError("DeleteDish", "DishStoragePort", err)
		return ErrInternal
	}

	return nil
}

func (cmd *DeleteDishCommand) belongsToRestaurant(ctx context.Context, restaurant uuid.UUID, dish uuid.UUID) (bool, error) {
	menus, err := cmd.menus.GetMenuByRestaurant(ctx, restaurant)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return false, ErrRestaurantNotFound
		}
		cmd.logInternalError("GetMenuByRestaurant", "MenuStoragePort", err)
		return false, ErrInternal
	}

	for _, m := range menus {
		if m.HasDish(dish) {
			return true, nil
		}
	}

	return false, nil
}

func (cmd *DeleteDishCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"DeleteDish Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewDeleteDishCommand(
	dishes DishStoragePort,
	restaurants RestaurantStoragePort,
	menus MenuStoragePort,
	logger *slog.Logger,
) *DeleteDishCommand {
	return &DeleteDishCommand{
		dishes:      dishes,
		restaurants: restaurants,
		menus:       menus,
		logger:      logger,
	}
}
