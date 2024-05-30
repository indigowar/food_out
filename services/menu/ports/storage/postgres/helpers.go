package postgres

import (
	"fmt"
	"net/url"

	"github.com/indigowar/food_out/services/menu/domain"
	"github.com/indigowar/food_out/services/menu/ports/storage/postgres/data"
)

// dishToModel converts data.Disht to domain.Dish
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
