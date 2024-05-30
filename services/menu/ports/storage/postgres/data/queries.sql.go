// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const retrieveDishByID = `-- name: RetrieveDishByID :one
SELECT id, name, image, price, is_deleted FROM dishes WHERE id = $1
`

func (q *Queries) RetrieveDishByID(ctx context.Context, id pgtype.UUID) (Dish, error) {
	row := q.db.QueryRow(ctx, retrieveDishByID, id)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Price,
		&i.IsDeleted,
	)
	return i, err
}

const retrieveDishByMenu = `-- name: RetrieveDishByMenu :many
SELECT d.id, d.name, d.image, d.price, d.is_deleted FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.id = $1
`

func (q *Queries) RetrieveDishByMenu(ctx context.Context, id pgtype.UUID) ([]Dish, error) {
	rows, err := q.db.Query(ctx, retrieveDishByMenu, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Price,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveDishesByRestaurant = `-- name: RetrieveDishesByRestaurant :many
SELECT d.id, d.name, d.image, d.price, d.is_deleted FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.restaurant = $1
`

func (q *Queries) RetrieveDishesByRestaurant(ctx context.Context, restaurant pgtype.UUID) ([]Dish, error) {
	rows, err := q.db.Query(ctx, retrieveDishesByRestaurant, restaurant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dish
	for rows.Next() {
		var i Dish
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Price,
			&i.IsDeleted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
