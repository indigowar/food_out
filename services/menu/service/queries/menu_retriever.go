package queries

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

//go:generate go run github.com/matryer/moq -out menu_retriever_moq_test.go . MenuRetriever

// MenuRetriever - a pore to the storage to retrieve the menus
type MenuRetriever interface {
	RetrieveByID(ctx context.Context, id uuid.UUID) (*domain.Menu, error)
	RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error)
}
