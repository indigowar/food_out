package commands

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type UpdateMenuArgument struct {
	ID         uuid.UUID
	Restaurant uuid.UUID
	Name       string
	Image      *url.URL
}

type UpdateMenuCommand struct {
	restaurant RestaurantStoragePort
	menu       MenuStoragePort

	logger *slog.Logger
}

func (cmd *UpdateMenuCommand) UpdateMenu(ctx context.Context, arg UpdateMenuArgument) error {
	if exists, err := cmd.restaurant.RestaurantExists(ctx, arg.Restaurant); !exists || err != nil {
		if err != nil {
			return err
		}
		return ErrRestaurantNotFound
	}

	menu, err := cmd.menu.GetMenu(ctx, arg.ID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrMenuNotFound
		}
		cmd.logInternalError("GetMenu", "MenuStoragePort", err)
		return ErrInternal
	}

	if menu.Restaurant() != arg.Restaurant {
		return ErrMenuNotFound
	}

	if err := cmd.updateMenu(menu, arg); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalidData, err.Error())
	}

	return cmd.saveMenu(ctx, menu)
}

func (cmd *UpdateMenuCommand) updateMenu(menu *domain.Menu, arg UpdateMenuArgument) error {
	if err := menu.SetName(arg.Name); err != nil {
		return err
	}

	if err := menu.SetImage(arg.Image); err != nil {
		return err
	}

	return nil
}

func (cmd *UpdateMenuCommand) saveMenu(ctx context.Context, menu *domain.Menu) error {
	if err := cmd.menu.UpdateMenu(ctx, menu); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrMenuNotFound
		}
		cmd.logInternalError("GetMenu", "MenuStoragePort", err)
		return ErrInternal
	}
	return nil
}

func (cmd *UpdateMenuCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"UpdateDish Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewUpdateMenuCommand(
	restaurant RestaurantStoragePort,
	menu MenuStoragePort,
	logger *slog.Logger,
) *UpdateMenuCommand {
	return &UpdateMenuCommand{
		restaurant: restaurant,
		menu:       menu,
		logger:     logger,
	}
}
