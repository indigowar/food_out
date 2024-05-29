package queries

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetDishesByMenuQuery struct {
	dish DishRetriever
	menu MenuRetriever

	logger *slog.Logger
}

func (q *GetDishesByMenuQuery) GetDishesByMenu(ctx context.Context, id uuid.UUID) ([]*domain.Dish, error) {
	if _, err := q.menu.RetrieveByID(ctx, id); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrMenuNotFound
		}
		q.logInternal("RetrieveByID", "MenuRetriever", err)
		return nil, ErrInternal
	}

	dishes, err := q.dish.RetrieveByMenu(ctx, id)
	if err != nil {
		q.logInternal("RetrieveByMenu", "DishRetriever", err)
		return nil, ErrInternal
	}

	return dishes, nil
}

func (q *GetDishesByMenuQuery) logInternal(action string, reason string, err error) {
	q.logger.Warn(
		"GetDishesByMenu Query FAILED",
		"action", action,
		"reason", reason,
		"error", err.Error(),
	)
}

func NewGetDishesByMenuQuery(
	dish DishRetriever,
	menu MenuRetriever,
	logger *slog.Logger,
) *GetDishesByMenuQuery {
	return &GetDishesByMenuQuery{
		dish:   dish,
		menu:   menu,
		logger: logger,
	}
}
