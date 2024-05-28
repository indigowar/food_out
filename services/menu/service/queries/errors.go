package queries

import (
	"errors"
	"fmt"
)

// Those errors are used both by the ports and by the queries
var (
	ErrNotFound    = errors.New("object is not found")
	ErrInvalidData = errors.New("invalid data provided")

	ErrDishNotFound       = fmt.Errorf("dish %w", ErrNotFound)
	ErrMenuNotFound       = fmt.Errorf("menu %w", ErrNotFound)
	ErrRestaurantNotFound = fmt.Errorf("restaurant %w", ErrNotFound)

	ErrInternal = errors.New("internal service error")
)
