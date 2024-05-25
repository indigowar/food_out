package queries

import "errors"

// Those errors are used both by the ports and by the queries
var (
	ErrDishNotFound       = errors.New("dish is not found")
	ErrMenuNotFound       = errors.New("menu is not found")
	ErrRestaurantNotFound = errors.New("restaurant is not found")
)
