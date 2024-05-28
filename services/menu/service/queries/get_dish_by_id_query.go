package queries

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetDishByIDQuery struct {
	retriever DishRetriever
	logger    *slog.Logger
}

func (q *GetDishByIDQuery) GetDishByID(ctx context.Context, id uuid.UUID) (*domain.Dish, error) {
	dish, err := q.retriever.RetrieveByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrDishNotFound
		}
		q.logger.Warn(
			"GetDishByID Query FAILED",
			"action", "RetrieveByID",
			"reason", "DishRetriever",
			"error", err.Error(),
		)
		return nil, ErrInternal
	}

	return dish, nil
}

func NewGetDishByIDQuery(
	retriever DishRetriever,
	logger *slog.Logger,
) *GetDishByIDQuery {
	return &GetDishByIDQuery{
		retriever: retriever,
		logger:    logger,
	}
}
