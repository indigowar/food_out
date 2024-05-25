package commands

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound      = errors.New("object is not found")
	ErrAlreadyExists = errors.New("object already exists")
	ErrInvalidData   = errors.New("invalid data provided")

	ErrDishNotFound      = fmt.Errorf("dish %w", ErrNotFound)
	ErrDishAlreadyExists = fmt.Errorf("dish %w", ErrAlreadyExists)

	ErrMenuNotFound      = fmt.Errorf("menu %w", ErrNotFound)
	ErrMenuAlreadyExists = fmt.Errorf("menu %w", ErrAlreadyExists)

	ErrRestaurantNotFound      = fmt.Errorf("restaurant %w", ErrNotFound)
	ErrRestaurantAlreadyExists = fmt.Errorf("restaurant %w", ErrAlreadyExists)
)
