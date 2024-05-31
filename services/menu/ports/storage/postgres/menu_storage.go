package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
	"github.com/indigowar/food_out/services/menu/service/commands"
)

type MenuStorage struct {
	conn *pgx.Conn
}

var _ commands.MenuStoragePort = &MenuStorage{}

// AddMenu implements commands.MenuStoragePort.
func (m *MenuStorage) AddMenu(ctx context.Context, menu *domain.Menu) error {
	return transaction(ctx, m.conn, func(ctx context.Context, queries *data.Queries) error {
		if err := queries.InsertMenu(ctx, data.InsertMenuParams{
			ID:         pgtype.UUID{Bytes: menu.ID(), Valid: true},
			Name:       menu.Name(),
			Image:      menu.Image().String(),
			Restaurant: pgtype.UUID{Bytes: menu.Restaurant(), Valid: true},
		}); err != nil {
			return err
		}

		for _, v := range menu.Dishes() {
			if err := queries.InsertIntoMenuDish(ctx, data.InsertIntoMenuDishParams{
				Menu: pgtype.UUID{Bytes: menu.ID(), Valid: true},
				Dish: pgtype.UUID{Bytes: v, Valid: true},
			}); err != nil {
				return err
			}
		}

		return nil
	})
}

// DeleteMenu implements commands.MenuStoragePort.
func (m *MenuStorage) DeleteMenu(ctx context.Context, id uuid.UUID) error {
	return transaction(ctx, m.conn, func(ctx context.Context, queries *data.Queries) error {
		if err := queries.DeleteMenu(ctx, pgtype.UUID{Bytes: id, Valid: true}); err != nil {
			return err
		}
		return nil
	})
}

// GetMenu implements commands.MenuStoragePort.
func (m *MenuStorage) GetMenu(ctx context.Context, id uuid.UUID) (*domain.Menu, error) {
	queries := data.New(m.conn)

	menu, err := queries.RetrieveMenuByID(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, commands.ErrMenuNotFound
		}
		return nil, err
	}

	dishes, err := queries.RetrieveDishesIdsByMenu(ctx, menu.ID)
	if err != nil {
		return nil, err
	}

	return menuToModel(menu, dishes)
}

// GetMenuByRestaurant implements commands.MenuStoragePort.
func (m *MenuStorage) GetMenuByRestaurant(ctx context.Context, restaurant uuid.UUID) ([]*domain.Menu, error) {
	queries := data.New(m.conn)
	ids, err := queries.RetrieveMenusIDsByRestaurant(ctx, pgtype.UUID{Bytes: restaurant, Valid: true})
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Menu, len(ids), 0)

	for _, v := range ids {
		menu, err := m.GetMenu(ctx, v.Bytes)
		if err != nil {
			return nil, err
		}
		result = append(result, menu)
	}

	return result, nil
}

// UpdateMenu implements commands.MenuStoragePort.
func (m *MenuStorage) UpdateMenu(ctx context.Context, menu *domain.Menu) error {
	return transaction(ctx, m.conn, func(ctx context.Context, queries *data.Queries) error {
		if err := queries.DeleteAllMenuDishByMenu(ctx, pgtype.UUID{Bytes: menu.ID(), Valid: true}); err != nil {
			return err
		}

		if err := queries.UpdateMenu(ctx, data.UpdateMenuParams{
			ID:         pgtype.UUID{Bytes: menu.ID(), Valid: true},
			Name:       menu.Name(),
			Image:      menu.Image().String(),
			Restaurant: pgtype.UUID{Bytes: menu.Restaurant(), Valid: true},
		}); err != nil {
			// todo: properly handle the error
			return err
		}

		for _, v := range menu.Dishes() {
			if err := queries.InsertIntoMenuDish(ctx, data.InsertIntoMenuDishParams{
				Menu: pgtype.UUID{Bytes: menu.ID(), Valid: true},
				Dish: pgtype.UUID{Bytes: v, Valid: true},
			}); err != nil {
				return err
			}
		}

		return nil
	})
}

func NewMenuStorage(conn *pgx.Conn) *MenuStorage {
	return &MenuStorage{
		conn: conn,
	}
}
