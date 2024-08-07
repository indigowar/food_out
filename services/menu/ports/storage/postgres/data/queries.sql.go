// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package data

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAllMenuDishByMenu = `-- name: DeleteAllMenuDishByMenu :exec
DELETE FROM menu_dish WHERE menu = $1
`

func (q *Queries) DeleteAllMenuDishByMenu(ctx context.Context, menu pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteAllMenuDishByMenu, menu)
	return err
}

const deleteDish = `-- name: DeleteDish :exec
DELETE FROM dishes
WHERE id = $1
`

func (q *Queries) DeleteDish(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteDish, id)
	return err
}

const deleteMenu = `-- name: DeleteMenu :exec
DELETE FROM menus
WHERE id = $1
`

func (q *Queries) DeleteMenu(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteMenu, id)
	return err
}

const deleteMenuDishByMenu = `-- name: DeleteMenuDishByMenu :exec
DELETE FROM menu_dish
WHERE menu = $1
`

func (q *Queries) DeleteMenuDishByMenu(ctx context.Context, menu pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteMenuDishByMenu, menu)
	return err
}

const getAllRestaurants = `-- name: GetAllRestaurants :many
SELECT id
FROM restaurants
`

func (q *Queries) GetAllRestaurants(ctx context.Context) ([]pgtype.UUID, error) {
	rows, err := q.db.Query(ctx, getAllRestaurants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.UUID
	for rows.Next() {
		var id pgtype.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRestaurant = `-- name: GetRestaurant :one
SELECT id
FROM restaurants
WHERE id = $1
`

func (q *Queries) GetRestaurant(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, getRestaurant, id)
	err := row.Scan(&id)
	return id, err
}

const insertDish = `-- name: InsertDish :exec
INSERT INTO dishes(id, name, image, price)
VALUES ($1, $2, $3, $4)
`

type InsertDishParams struct {
	ID    pgtype.UUID
	Name  string
	Image string
	Price float64
}

func (q *Queries) InsertDish(ctx context.Context, arg InsertDishParams) error {
	_, err := q.db.Exec(ctx, insertDish,
		arg.ID,
		arg.Name,
		arg.Image,
		arg.Price,
	)
	return err
}

const insertIntoMenuDish = `-- name: InsertIntoMenuDish :exec
INSERT INTO menu_dish(menu, dish)
VALUES ($1, $2)
`

type InsertIntoMenuDishParams struct {
	Menu pgtype.UUID
	Dish pgtype.UUID
}

func (q *Queries) InsertIntoMenuDish(ctx context.Context, arg InsertIntoMenuDishParams) error {
	_, err := q.db.Exec(ctx, insertIntoMenuDish, arg.Menu, arg.Dish)
	return err
}

const insertMenu = `-- name: InsertMenu :exec
INSERT INTO menus(id, name, image, restaurant)
VALUES ($1, $2, $3, $4)
`

type InsertMenuParams struct {
	ID         pgtype.UUID
	Name       string
	Image      string
	Restaurant pgtype.UUID
}

func (q *Queries) InsertMenu(ctx context.Context, arg InsertMenuParams) error {
	_, err := q.db.Exec(ctx, insertMenu,
		arg.ID,
		arg.Name,
		arg.Image,
		arg.Restaurant,
	)
	return err
}

const insertRestaurant = `-- name: InsertRestaurant :exec
INSERT INTO restaurants(id)
VALUES ($1)
`

func (q *Queries) InsertRestaurant(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, insertRestaurant, id)
	return err
}

const retrieveDishByID = `-- name: RetrieveDishByID :one
SELECT id, name, image, price FROM dishes WHERE id = $1
`

func (q *Queries) RetrieveDishByID(ctx context.Context, id pgtype.UUID) (Dish, error) {
	row := q.db.QueryRow(ctx, retrieveDishByID, id)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Price,
	)
	return i, err
}

const retrieveDishByMenu = `-- name: RetrieveDishByMenu :many
SELECT d.id, d.name, d.image, d.price FROM dishes d
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
SELECT d.id, d.name, d.image, d.price FROM dishes d
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

const retrieveDishesIdsByMenu = `-- name: RetrieveDishesIdsByMenu :many
SELECT md.dish FROM menu_dish md
JOIN menus m ON md.menu = m.id
WHERE menu = $1
`

func (q *Queries) RetrieveDishesIdsByMenu(ctx context.Context, menu pgtype.UUID) ([]pgtype.UUID, error) {
	rows, err := q.db.Query(ctx, retrieveDishesIdsByMenu, menu)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.UUID
	for rows.Next() {
		var dish pgtype.UUID
		if err := rows.Scan(&dish); err != nil {
			return nil, err
		}
		items = append(items, dish)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const retrieveMenuByID = `-- name: RetrieveMenuByID :one
SELECT id, name, image, restaurant FROM menus WHERE id = $1
`

func (q *Queries) RetrieveMenuByID(ctx context.Context, id pgtype.UUID) (Menu, error) {
	row := q.db.QueryRow(ctx, retrieveMenuByID, id)
	var i Menu
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Restaurant,
	)
	return i, err
}

const retrieveMenusByRestaurant = `-- name: RetrieveMenusByRestaurant :many
SELECT id, name, image, restaurant FROM menus WHERE restaurant = $1
`

func (q *Queries) RetrieveMenusByRestaurant(ctx context.Context, restaurant pgtype.UUID) ([]Menu, error) {
	rows, err := q.db.Query(ctx, retrieveMenusByRestaurant, restaurant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Menu
	for rows.Next() {
		var i Menu
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
			&i.Restaurant,
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

const retrieveMenusIDsByRestaurant = `-- name: RetrieveMenusIDsByRestaurant :many
SELECT id
FROM menus
WHERE restaurant = $1
`

func (q *Queries) RetrieveMenusIDsByRestaurant(ctx context.Context, restaurant pgtype.UUID) ([]pgtype.UUID, error) {
	rows, err := q.db.Query(ctx, retrieveMenusIDsByRestaurant, restaurant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.UUID
	for rows.Next() {
		var id pgtype.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDish = `-- name: UpdateDish :exec
UPDATE dishes
SET name = $2, image = $3, price = $4
WHERE id = $1
`

type UpdateDishParams struct {
	ID    pgtype.UUID
	Name  string
	Image string
	Price float64
}

func (q *Queries) UpdateDish(ctx context.Context, arg UpdateDishParams) error {
	_, err := q.db.Exec(ctx, updateDish,
		arg.ID,
		arg.Name,
		arg.Image,
		arg.Price,
	)
	return err
}

const updateMenu = `-- name: UpdateMenu :exec
UPDATE menus
SET name = $2, image = $3, restaurant = $4
WHERE id = $1
`

type UpdateMenuParams struct {
	ID         pgtype.UUID
	Name       string
	Image      string
	Restaurant pgtype.UUID
}

func (q *Queries) UpdateMenu(ctx context.Context, arg UpdateMenuParams) error {
	_, err := q.db.Exec(ctx, updateMenu,
		arg.ID,
		arg.Name,
		arg.Image,
		arg.Restaurant,
	)
	return err
}
