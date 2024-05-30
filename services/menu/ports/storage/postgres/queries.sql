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
