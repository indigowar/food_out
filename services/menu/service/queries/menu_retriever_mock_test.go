package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

// mockMenuRetrieveByID - mocks MenuRetriver.RetrieveByID
func mockMenuRetrieveByID(menu *domain.Menu, e error) func(context.Context, uuid.UUID) (*domain.Menu, error) {
	return func(_ context.Context, _ uuid.UUID) (*domain.Menu, error) {
		return menu, e
	}
}

// RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error)
func mockMenuRetrieveByRestaurant(menus []*domain.Menu, err error) func(context.Context, uuid.UUID) ([]*domain.Menu, error) {
	return func(_ context.Context, _ uuid.UUID) ([]*domain.Menu, error) {
		return menus, err
	}
}
