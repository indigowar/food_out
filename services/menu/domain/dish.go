package domain

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

type Dish struct {
	ID    uuid.UUID
	Name  string
	Image *url.URL
	Price float64
}

func NewDish(name string, image *url.URL, price float64) (*Dish, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	if image == nil {
		return nil, errors.New("invalid image url")
	}

	if price <= 0 {
		return nil, errors.New("price is invalid")
	}

	return &Dish{
		ID:    uuid.New(),
		Name:  name,
		Image: image,
		Price: price,
	}, nil
}
