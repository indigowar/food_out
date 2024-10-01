package models

import "github.com/google/uuid"

type OriginalProduct struct {
	ID          uuid.UUID
	Restaurant  uuid.UUID
	Name        string
	Picture     string
	Price       float64
	Description string
}

type Product struct {
	ID          uuid.UUID
	Original    uuid.UUID
	Restaurant  uuid.UUID
	Name        string
	Picture     string
	Price       float64
	Description string
}
