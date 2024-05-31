// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package data

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Dish struct {
	ID    pgtype.UUID
	Name  string
	Image string
	Price float64
}

type Menu struct {
	ID         pgtype.UUID
	Name       string
	Image      string
	Restaurant pgtype.UUID
}

type MenuDish struct {
	Menu pgtype.UUID
	Dish pgtype.UUID
}

type Restaurant struct {
	ID pgtype.UUID
}
