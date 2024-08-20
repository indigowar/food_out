package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Restaurant  uuid.UUID
	Name        string
	Picture     string
	Price       float64
	Description string
	Categories  []uuid.UUID
}
