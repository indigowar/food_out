-- name: GetOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: GetOrderAcceptanceByID :one
SELECT * FROM orders_acceptance WHERE order_id = $1;

-- name: GetOrderCourierByID :one
SELECT * FROM orders_couriers WHERE order_id = $1;

-- name: GetOrderPaymentByID :one
SELECT * FROM orders_payments WHERE order_id = $1;

-- name: GetProductForOrder :many
SELECT * FROM products WHERE order_id = $1;

-- name: InsertProduct :execresult
INSERT INTO products (
    id, original_id, restaurant, order_id,
    name, picture, price, description
) VALUES (
    $1, $2, $3, $4,
    $5, $6, $7, $8
);

-- name: GetProductCategories :many
SELECT category FROM products_categories WHERE product = $1;

-- name: InsertProductCategory :execresult
INSERT INTO products_categories(product, category)
VALUES ($1, $2);
