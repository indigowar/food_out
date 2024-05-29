package queries

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type GetRestaurantsQuery struct {
	retriever RestaurantRetriever
	logger    *slog.Logger
}

func (q *GetRestaurantsQuery) GetRestaurants(ctx context.Context) ([]uuid.UUID, error) {
	ids, err := q.retriever.GetRestaurants(ctx)
	if err != nil {
		q.logger.Warn(
			"GetRestaurants Query FAILED",
			"action", "GetRestaurants",
			"reason", "RestaurantRetriever",
			"error", err.Error(),
		)
	}
	return ids, nil
}

func NewGetRestaurantsQuery(
	retriever RestaurantRetriever,
	logger *slog.Logger,
) *GetRestaurantsQuery {
	return &GetRestaurantsQuery{
		retriever: retriever,
		logger:    logger,
	}
}
