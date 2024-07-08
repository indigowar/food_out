package queries

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

type GetMenuByIDQuery struct {
	logger    *slog.Logger
	retriever MenuRetriever
}

func (q *GetMenuByIDQuery) GetMenuByID(ctx context.Context, id uuid.UUID) (*domain.Menu, error) {
	menu, err := q.retriever.RetrieveByID(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, ErrMenuNotFound
		}
		return nil, ErrInternal
	}
	return menu, nil
}

func NewGetMenuByIdQuery(
	logger *slog.Logger,
	retriever MenuRetriever,
) *GetMenuByIDQuery {
	return &GetMenuByIDQuery{
		logger:    logger,
		retriever: retriever,
	}
}
