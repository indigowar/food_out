package queries

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetRestaurantsMenusQuery struct {
	menu       MenuRetriever
	restaurant RestaurantRetriever

	logger *slog.Logger
}

func (q *GetRestaurantsMenusQuery) GetRestaurantsMenus(ctx context.Context, id uuid.UUID) ([]*domain.Menu, error) {
	if exists, err := q.restaurant.RestaurantExists(ctx, id); !exists || err != nil {
		if err != nil {
			q.logInternal("RestaurantExists", "RestaurantRetriever", err)
			return nil, ErrInternal
		}
		return nil, ErrRestaurantNotFound
	}

	menus, err := q.menu.RetrieveByRestaurant(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrMenuNotFound
		}
		q.logInternal("RestaurantExists", "RestaurantRetriever", err)
		return nil, ErrInternal
	}
	return menus, nil
}

func (q *GetRestaurantsMenusQuery) logInternal(action string, reason string, err error) {
	q.logger.Warn(
		"GetRestaurantsMenus Query FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewGetRestaurantsMenuQuery(
	menu MenuRetriever,
	restaurant RestaurantRetriever,
	logger *slog.Logger,
) *GetRestaurantsMenusQuery {
	return &GetRestaurantsMenusQuery{
		menu:       menu,
		restaurant: restaurant,
		logger:     logger,
	}
}
