package domain

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

type Menu struct {
	id         uuid.UUID
	name       string
	restaurant uuid.UUID
	image      *url.URL
	dishes     map[uuid.UUID]struct{}
}

func (m *Menu) ID() uuid.UUID {
	return m.id
}

func (m *Menu) Name() string {
	return m.name
}

func (m *Menu) SetName(value string) error {
	if len(value) == 0 {
		return errors.New("name is empty")
	}
	m.name = value
	return nil
}

func (m *Menu) Image() *url.URL {
	return m.image
}

func (m *Menu) SetImage(value *url.URL) error {
	if value == nil {
		return errors.New("image is empty")
	}
	m.image = value
	return nil
}

func (m *Menu) Dishes() []uuid.UUID {
	dishes := make([]uuid.UUID, len(m.dishes), 0)

	i := 0
	for v, _ := range m.dishes {
		dishes[i] = v
	}

	return dishes
}

func (m *Menu) AddDish(dish uuid.UUID) {
	m.dishes[dish] = struct{}{}
}

func (m *Menu) RemoveDish(dish uuid.UUID) error {
	if _, ok := m.dishes[dish]; !ok {
		return errors.New("dish is not in menu")
	}
	delete(m.dishes, dish)
	return nil
}

func NewMenu(name string, restaurant uuid.UUID, image *url.URL) (*Menu, error) {
	menu := Menu{id: uuid.New(), restaurant: restaurant, dishes: make(map[uuid.UUID]struct{})}

	if err := menu.SetName(name); err != nil {
		return nil, err
	}

	if err := menu.SetImage(image); err != nil {
		return nil, err
	}

	return &menu, nil
}
