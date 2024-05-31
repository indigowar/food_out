-- name: RetrieveDishByID :one
SELECT * FROM dishes WHERE id = $1 AND is_deleted = FALSE;

-- name: RetrieveDishByMenu :many
SELECT d.* FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.id = $1 AND m.is_deleted = FALSE AND d.is_deleted = FALSE;

-- name: RetrieveDishesByRestaurant :many
SELECT d.* FROM dishes d
JOIN menu_dish md ON dm.dish = d.id
JOIN menus m ON m.id = md.menu
WHERE m.restaurant = $1 AND m.is_deleted = FALSE;

-- name: RetrieveDishesIdsByMenu :many
SELECT dish FROM menu_dish
WHERE menu = $1;

-- name: RetrieveMenuByID :one
SELECT * FROM menus WHERE id = $1 AND is_deleted = FALSE;

-- name: RetrieveMenusByRestaurant :many
SELECT * FROM menus WHERE restaurant = $1 AND is_deleted = FALSE;

-- name: InsertDish :exec
INSERT INTO dishes(id, name, image, price)
VALUES ($1, $2, $3, $4);

-- name: DeleteDish :exec
UPDATE dishes
SET is_deleted = TRUE
WHERE id = $1 AND is_deleted = FALSE;

-- name: UpdateDish :exec
UPDATE dishes
SET name = $2, image = $3, price = $4
WHERE id = $1 AND is_deleted = FALSE;