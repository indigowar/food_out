package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/queries"
)

type MenuRetriever struct {
	queries *data.Queries
}

var _ queries.MenuRetriever = &MenuRetriever{}

// RetrieveByID implements queries.MenuRetriever.
func (m *MenuRetriever) RetrieveByID(ctx context.Context, id uuid.UUID) (*domain.Menu, error) {
	result, err := m.queries.RetrieveMenuByID(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, queries.ErrNotFound
		}
		return nil, err
	}

	dishes, err := m.queries.RetrieveDishesIdsByMenu(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			dishes = make([]pgtype.UUID, 0)
		} else {
			return nil, err
		}
	}

	return menuToModel(result, dishes)
}

// RetrieveByRestaurant implements queries.MenuRetriever.
func (m *MenuRetriever) RetrieveByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error) {
	menus, err := m.queries.RetrieveMenusByRestaurant(ctx, pgtype.UUID{Bytes: restaurant, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, queries.ErrNotFound
		}
		return nil, err
	}

	result := make([]*domain.Menu, len(menus), 0)

	for _, v := range menus {
		dishes, err := m.queries.RetrieveDishesIdsByMenu(ctx, v.ID)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				dishes = make([]pgtype.UUID, 0)
			} else {
				return nil, err
			}
		}

		r, err := menuToModel(v, dishes)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}

	return result, nil
}

func NewMenuRetriever(conn *pgx.Conn) *MenuRetriever {
	return &MenuRetriever{
		queries: data.New(conn),
	}
}
