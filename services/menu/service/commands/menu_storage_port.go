package commands

import (
	"context"

	"github.com/google/uuid"

	"github.com/indigowar/food_out/services/menu/domain"
)

//go:generate go run github.com/matryer/moq -out menu_storage_port_moq_test.go . MenuStoragePort

// MenuStoragePort - is a port to menu storage
type MenuStoragePort interface {
	GetMenuByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error)
	GetMenu(ctx context.Context, id uuid.UUID) (*domain.Menu, error)
	AddMenu(ctx context.Context, menu *domain.Menu) error
	DeleteMenu(ctx context.Context, id uuid.UUID) error
	UpdateMenu(ctx context.Context, menu *domain.Menu) error
}
