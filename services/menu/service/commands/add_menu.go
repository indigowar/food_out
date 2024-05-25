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
	menu, err := domain.NewMenu(args.Name, args.Restaurant, args.Image)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	empty := uuid.UUID{}
	if args.Restaurant == empty {
		return uuid.UUID{}, ErrInvalidData
	}

	if exists, err := cmd.restaurants.RestaurantExists(ctx, args.Restaurant); !exists || err != nil {
		if !exists {
			return uuid.UUID{}, ErrRestaurantNotFound
		}
		cmd.logInternalError("RestaurantExists", "RestaurantStoragePort", err)
		return uuid.UUID{}, err
	}

	if err := cmd.menus.AddMenu(ctx, menu); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return uuid.UUID{}, ErrMenuAlreadyExists
		}
		cmd.logInternalError("AddMenu", "MenuStoragePort", err)
		return uuid.UUID{}, err
	}

	return menu.ID(), nil
}

func (cmd *AddMenuCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"AddMenu Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
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
