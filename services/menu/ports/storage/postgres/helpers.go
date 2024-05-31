package postgres

import (
	"fmt"
	"net/url"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
)

// dishToModel converts data.Dish to domain.Dish
func dishToModel(data data.Dish) (*domain.Dish, error) {
	imageUrl, err := url.Parse(data.Image)
	if err != nil {
		return nil, fmt.Errorf("storage contains invalid data: %w", err)
	}

	dish, err := domain.NewDishWithID(
		data.ID.Bytes,
		data.Name,
		imageUrl,
		data.Price,
	)

	if err != nil {
		return nil, fmt.Errorf("storage contains invalid data: %w", err)
	}

	return dish, nil
}

// menuToModel - converts data.Menu and []pgtype.UUID to domain.Menu
func menuToModel(menu data.Menu, dishes []pgtype.UUID) (*domain.Menu, error) {
	imageUrl, err := url.Parse(menu.Image)
	if err != nil {
		return nil, fmt.Errorf("storage contains invalid data: %w", err)
	}

	dishSet := make(map[uuid.UUID]struct{}, len(dishes))
	for _, v := range dishes {
		dishSet[v.Bytes] = struct{}{}
	}

	return domain.NewMenuFull(
		menu.ID.Bytes,
		menu.Name,
		menu.Restaurant.Bytes,
		imageUrl,
		dishSet,
	)
}

func isDuplicatedKeyError(err error) bool {
	if pgError, ok := (err).(*pgconn.PgError); ok && pgError.Code == "23505" {
		return true
	}
	return false
}
