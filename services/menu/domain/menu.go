package domain

import (
	"net/url"

	"github.com/google/uuid"
)

type Menu struct {
	ID         uuid.UUID
	Name       string
	Restaurant uuid.UUID
	Image      *url.URL
	Dishes     []uuid.UUID
}

func (m *Menu) AddDish(id uuid.UUID) {
	m.Dishes = append(m.Dishes, id)
}
