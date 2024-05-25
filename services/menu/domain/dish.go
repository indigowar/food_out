package domain

import (
	"net/url"

	"github.com/google/uuid"
)

type Dish struct {
	ID    uuid.UUID
	Name  string
	Image *url.URL
	Price float64
}
