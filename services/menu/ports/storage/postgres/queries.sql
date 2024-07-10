-- name: RetrieveDishByID :one
SELECT * FROM dishes WHERE id = $1;

-- name: RetrieveDishByMenu :many
SELECT d.* FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.id = $1;

-- name: RetrieveDishesByRestaurant :many
SELECT d.* FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.restaurant = $1;

-- name: RetrieveDishesIdsByMenu :many
SELECT md.dish FROM menu_dish md
JOIN menus m ON md.menu = m.id
WHERE menu = $1;

-- name: RetrieveMenuByID :one
SELECT * FROM menus WHERE id = $1;

-- name: RetrieveMenusByRestaurant :many
SELECT * FROM menus WHERE restaurant = $1;

-- name: InsertDish :exec
INSERT INTO dishes(id, name, image, price)
VALUES ($1, $2, $3, $4);

-- name: DeleteDish :exec
DELETE FROM dishes
WHERE id = $1;

-- name: UpdateDish :exec
UPDATE dishes
SET name = $2, image = $3, price = $4
WHERE id = $1;

-- name: InsertMenu :exec
INSERT INTO menus(id, name, image, restaurant)
VALUES ($1, $2, $3, $4);

-- name: DeleteMenuDishByMenu :exec
DELETE FROM menu_dish
WHERE menu = $1;

-- name: InsertIntoMenuDish :exec
INSERT INTO menu_dish(menu, dish)
VALUES ($1, $2);

-- name: DeleteMenu :exec
DELETE FROM menus
WHERE id = $1;

-- name: DeleteAllMenuDishByMenu :exec
DELETE FROM menu_dish WHERE menu = $1;

-- name: UpdateMenu :exec
UPDATE menus
SET name = $2, image = $3, restaurant = $4
WHERE id = $1;

-- name: RetrieveMenusIDsByRestaurant :many
SELECT id
FROM menus
WHERE restaurant = $1;

-- name: InsertRestaurant :exec
INSERT INTO restaurants(id)
VALUES ($1);

-- name: GetRestaurant :one
SELECT *
FROM restaurants
WHERE id = $1;

-- name: GetAllRestaurants :many
SELECT *
FROM restaurants;