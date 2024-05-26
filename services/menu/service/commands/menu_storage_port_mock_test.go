package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

func mockGetMenu(menu *domain.Menu, err error) func(context.Context, uuid.UUID) (*domain.Menu, error) {
	return func(_ context.Context, _ uuid.UUID) (*domain.Menu, error) {
		return menu, err
	}
}

func mockGetMenuByRestaurant(menu []*domain.Menu, err error) func(context.Context, uuid.UUID) ([]*domain.Menu, error) {
	return func(_ context.Context, _ uuid.UUID) ([]*domain.Menu, error) {
		return menu, err
	}
}

func mockAddMenu(err error) func(context.Context, *domain.Menu) error {
	return func(_ context.Context, _ *domain.Menu) error {
		return err
	}
}

func mockDeleteMenu(err error) func(context.Context, uuid.UUID) error {
	return func(_ context.Context, _ uuid.UUID) error {
		return err
	}
}

func mockUpdateMenu(err error) func(context.Context, *domain.Menu) error {
	return func(_ context.Context, _ *domain.Menu) error {
		return err
	}
}
