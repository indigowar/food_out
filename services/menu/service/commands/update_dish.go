package commands

import (
	"context"
	"errors"
	"log/slog"
	"net/url"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type UpdateDishArgument struct {
	Id         uuid.UUID
	Restaurant uuid.UUID
	Name       string
	Image      *url.URL
	Price      float64
}

type UpdateDishCommand struct {
	restaurants RestaurantStoragePort
	menu        MenuStoragePort
	dishes      DishStoragePort

	logger *slog.Logger
}

func (cmd *UpdateDishCommand) UpdateDish(ctx context.Context, arg UpdateDishArgument) error {
	if !cmd.validateArgument(arg) {
		return ErrInvalidData
	}

	if err := cmd.checkRestaurantExistence(ctx, arg.Restaurant); err != nil {
		return err
	}

	if belongs, err := cmd.dishBelongsToRestaurant(ctx, arg.Restaurant, arg.Id); !belongs || err != nil {
		if err != nil {
			return err
		}

		return ErrDishNotFound
	}

	dish, err := cmd.findDish(ctx, arg.Id)
	if err != nil {
		return err
	}

	if err := cmd.updateDish(dish, arg); err != nil {
		return err
	}

	return cmd.saveDish(ctx, dish)
}

func (cmd *UpdateDishCommand) validateArgument(arg UpdateDishArgument) bool {
	emptyId := uuid.UUID{}
	return arg.Id == emptyId || arg.Restaurant == emptyId
}

func (cmd *UpdateDishCommand) checkRestaurantExistence(ctx context.Context, id uuid.UUID) error {
	exists, err := cmd.restaurants.RestaurantExists(ctx, id)
	if err != nil {
		cmd.logInternalError("RestaurantExists", "RestaurantStoragePort", err)
		return ErrInternal
	}

	if !exists {
		return ErrRestaurantNotFound
	}

	return nil
}

func (cmd *UpdateDishCommand) dishBelongsToRestaurant(ctx context.Context, restaurant uuid.UUID, dish uuid.UUID) (bool, error) {
	menus, err := cmd.menu.GetMenuByRestaurant(ctx, restaurant)
	if err != nil {
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

func (cmd *UpdateDishCommand) findDish(ctx context.Context, id uuid.UUID) (*domain.Dish, error) {
	dish, err := cmd.dishes.GetDish(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrDishNotFound
		}

		cmd.logInternalError("GetDish", "DishStoragePort", err)
		return nil, ErrInternal
	}

	return dish, nil
}

func (cmd *UpdateDishCommand) updateDish(dish *domain.Dish, arg UpdateDishArgument) error {
	if err := dish.SetName(arg.Name); err != nil {
		return err
	}

	if err := dish.SetImage(arg.Image); err != nil {
		return err
	}

	if err := dish.SetPrice(arg.Price); err != nil {
		return err
	}

	return nil
}

func (cmd *UpdateDishCommand) saveDish(ctx context.Context, dish *domain.Dish) error {
	if err := cmd.dishes.UpdateDish(ctx, dish); err != nil {
		if errors.Is(err, ErrNotFound) {
			return ErrDishNotFound
		}

		cmd.logInternalError("UpdateDish", "DishStoragePort", err)
		return ErrInternal
	}

	return nil
}

func (cmd *UpdateDishCommand) logInternalError(action string, reason string, err error) {
	cmd.logger.Warn(
		"UpdateDish Command FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewUpdateDishCommand(
	restaurants RestaurantStoragePort,
	menu MenuStoragePort,
	dishes DishStoragePort,
	logger *slog.Logger,
) *UpdateDishCommand {
	return &UpdateDishCommand{
		restaurants: restaurants,
		menu:        menu,
		dishes:      dishes,
		logger:      logger,
	}
}
