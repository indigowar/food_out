package commands

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type DeleteMenuCommand struct {
	restaurants RestaurantStoragePort
	menus       MenuStoragePort
	logger      *slog.Logger
}

func (cmd *DeleteMenuCommand) DeleteMenu(ctx context.Context, restaurant uuid.UUID, menuId uuid.UUID) error {
	empty := uuid.UUID{}
	if restaurant == empty || menuId == empty {
		return ErrInvalidData
	}

	if exists, err := cmd.restaurants.RestaurantExists(ctx, restaurant); err != nil || !exists {
		if err != nil {
			cmd.logInternalError("RestaurantExists", "RestaurantStoragePort", err)
			return ErrInternal
		}
		return ErrRestaurantNotFound
	}

	menu, err := cmd.menus.GetMenu(ctx, menuId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrMenuNotFound
		}
		cmd.logInternalError("GetMenu", "MenuStoragePort", err)
		return ErrInternal
	}

	if menu.Restaurant() != restaurant {
		return ErrMenuNotFound
	}

	if err := cmd.menus.DeleteMenu(ctx, menu.ID()); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrMenuNotFound
		}
		cmd.logInternalError("GetMenu", "MenuStoragePort", err)
		return ErrInternal
	}

	return nil
}

func (cmd *DeleteMenuCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"DeleteMenu Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewDeleteMenuCommand(
	restaurants RestaurantStoragePort,
	menus MenuStoragePort,
	logger *slog.Logger,
) *DeleteMenuCommand {
	return &DeleteMenuCommand{
		restaurants: restaurants,
		menus:       menus,
		logger:      logger,
	}
}
