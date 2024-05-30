package queries

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type ValidateDishOwnershipQuery struct {
	menus      MenuRetriever
	restaurant RestaurantRetriever

	logger *slog.Logger
}

func (q *ValidateDishOwnershipQuery) ValidateDishOwnership(ctx context.Context, restaurantId uuid.UUID, dishesId []uuid.UUID) (bool, error) {
	if exists, err := q.restaurant.RestaurantExists(ctx, restaurantId); !exists || err != nil {
		if err != nil {
			q.logInternalError("RestaurantExists", "RestaurantRetriever", err)
			return false, ErrInternal
		}
		return false, ErrRestaurantNotFound
	}

	menus, err := q.menus.RetrieveByRestaurant(ctx, restaurantId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return false, ErrRestaurantNotFound
		}
		q.logInternalError("RetrieveByRestaurant", "MenuRetriever", err)
		return false, ErrInternal
	}

	dishPresence := make(map[uuid.UUID]bool)

	for _, dish := range dishesId {
		dishPresence[dish] = false
		for _, menu := range menus {
			if menu.HasDish(dish) {
				dishPresence[dish] = true
			}
		}
	}

	for _, precense := range dishPresence {
		if !precense {
			return false, nil
		}
	}

	return true, nil
}

func (q *ValidateDishOwnershipQuery) logInternalError(action string, reason string, err error) {
	q.logger.Warn(
		"ValidateDishOwnership Query FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewValidateDishOwnershipQuery(
	menus MenuRetriever,
	restaurant RestaurantRetriever,
	logger *slog.Logger,
) *ValidateDishOwnershipQuery {
	return &ValidateDishOwnershipQuery{
		menus:      menus,
		restaurant: restaurant,
		logger:     logger,
	}
}
