-- name: SelectProductCategoriesByProductID :many
SELECT * FROM products_categories WHERE product = $1;

-- name: SelectProductsByOrderID :many
SELECT * FROM products WHERE order_id = $1;

-- name: SelectProductsIDByOrderID :many
SELECT id FROM products WHERE order_id = $1;

-- name: InsertOrUpdateProduct :execresult
INSERT INTO products(
    id, original_id, restaurant, order_id, name, picture, price, description
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) ON CONFLICT (id) DO UPDATE
SET original_id = EXCLUDED.original_id,
    restaurant = EXCLUDED.restaurant,
    order_id = EXCLUDED.order_id,
    name = EXCLUDED.name,
    picture = EXCLUDED.picture,
    price = EXCLUDED.price,
    description = EXCLUDED.description;

-- name: DeleteProductByID :exec
DELETE FROM products WHERE id = $1;
