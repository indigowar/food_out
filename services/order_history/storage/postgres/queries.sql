-- name: InsertOrder :exec
INSERT INTO orders(
    id,
    restaurant,
    created_at,
    customer,
    customer_address,
    cooking_started_at,
    delivery_started_at,
    delivery_completed_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: InsertOrdersAcceptance :exec
INSERT INTO orders_acceptance(order_id, manager, accepted_at)
VALUES ($1, $2, $3);

-- name: InsertOrdersCourier :exec
INSERT INTO orders_couriers(order_id, courier, taken_at)
VALUES ($1, $2, $3);

-- name: InsertOrdersPayment :exec
INSERT INTO orders_payments(order_id, transaction, payed_at)
VALUES ($1, $2, $3);

-- name: InsertOrdersCancellation :exec
INSERT INTO orders_cancellations(order_id, canceller, cancelled_at)
VALUES ($1, $2, $3);

