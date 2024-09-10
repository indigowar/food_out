-- name: SelectOrderByID :one
SELECT * FROM orders WHERE id = $1;

-- name: SelectOrderAcceptanceByID :one
SELECT * FROM orders_acceptance WHERE order_id = $1;

-- name: SelectOrdersCourierByOrderID :one
SELECT * FROM orders_couriers WHERE order_id = $1;

-- name: SelectOrdersPaymentByOrderID :one
SELECT * FROM orders_payments WHERE order_id = $1;

-- name: SelectOrdersCancellationByOrderID :one
SELECT * FROM orders_cancellations WHERE order_id = $1;
