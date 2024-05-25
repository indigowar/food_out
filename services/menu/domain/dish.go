package domain

import (
	"errors"
	"net/url"

	"github.com/google/uuid"
)

type Dish struct {
	id    uuid.UUID
	name  string
	image *url.URL
	price float64
}

func (d *Dish) ID() uuid.UUID {
	return d.id
}

func (d *Dish) Name() string {
	return d.name
}

func (d *Dish) SetName(value string) error {
	if len(value) == 0 {
		return errors.New("name is empty")
	}
	d.name = value
	return nil
}

func (d *Dish) Image() *url.URL {
	return d.image
}

func (d *Dish) SetImage(value *url.URL) error {
	if value == nil {
		return errors.New("image is empty")
	}
	d.image = value
	return nil
}

func (d *Dish) SetPrice(value float64) error {
	if value <= 0 {
		return errors.New("price should be positive")
	}
	d.price = value
	return nil
}

func (d *Dish) Price() float64 {
	return d.price
}

func NewDish(name string, image *url.URL, price float64) (*Dish, error) {
	dish := Dish{id: uuid.New()}

	if err := dish.SetName(name); err != nil {
		return nil, err
	}

	if err := dish.SetImage(image); err != nil {
		return nil, err
	}

	if err := dish.SetPrice(price); err != nil {
		return nil, err
	}

	return &dish, nil
}
