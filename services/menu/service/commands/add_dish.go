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
	dish, err := domain.NewDish(args.Name, args.Image, args.Price)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w: %s", ErrInvalidData, err)
	}

	emptyId := uuid.UUID{}
	if args.Restaurant == emptyId {
		return uuid.UUID{}, fmt.Errorf("%w: restaurant id", ErrInvalidData)
	}
	if args.Menu == emptyId {
		return uuid.UUID{}, fmt.Errorf("%w: menu id", ErrInvalidData)
	}

	if exists, err := cmd.restaurants.RestaurantExists(ctx, args.Restaurant); !exists || err != nil {
		if !exists {
			return uuid.UUID{}, ErrRestaurantNotFound
		}

		cmd.logInternalError("RestaurantExists", "RestaurantStoragePort", err)
		return uuid.UUID{}, err
	}

	menu, err := cmd.menus.GetMenu(ctx, args.Menu)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return uuid.UUID{}, ErrMenuNotFound
		}
		cmd.logInternalError("GetMenu", "MenuStoragePort", err)
		return uuid.UUID{}, err
	}

	menu.AddDish(dish.ID)

	if err := cmd.dishes.AddDish(ctx, dish); err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return uuid.UUID{}, ErrDishAlreadyExists
		}
		cmd.logInternalError("AddDish", "DishStoragePort", err)
		return uuid.UUID{}, err
	}

	if err := cmd.menus.UpdateMenu(ctx, menu); err != nil {
		cmd.logInternalError("UpdateMenu", "MenuStoragePort", err)
		return uuid.UUID{}, err
	}

	return dish.ID, nil
}

func (cmd *AddDishCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"AddDish Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
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
