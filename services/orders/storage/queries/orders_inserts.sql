-- name: InsertOrder :execresult
INSERT INTO orders(id, restaurant, created_at, customer, customer_address)
VALUES ($1, $2, $3, $4, $5);

-- name: InserOrUpdateOrdersAcceptance :execresult
INSERT INTO orders_acceptance(order_id, manager, accepted_at)
VALUES ($1, $2, $3)
ON CONFLICT (order_id) DO UPDATE
SET manager = EXCLUDED.manager,
    accepted_at = EXCLUDED.accepted_at;

-- name: InsertOrUpdateOrdersCourier :execresult
INSERT INTO orders_couriers(order_id, courier, taken_at)
VALUES ($1, $2, $3)
ON CONFLICT (order_id) DO UPDATE
SET courier = EXCLUDED.courier,
    taken_at = EXCLUDED.taken_at;

-- name: InsertOrUpdateOrdersPayment :execresult
INSERT INTO orders_payments(order_id, transaction, payed_at)
VALUES ($1, $2, $3)
ON CONFLICT (order_id) DO UPDATE
SET transaction = EXCLUDED.transaction,
    payed_at = EXCLUDED.payed_at;

-- name: InsertOrUpdateOrdersCancellation :execresult
INSERT INTO orders_cancellations(order_id, canceller, cancelled_at)
VALUES ($1, $2, $3)
ON CONFLICT (order_id) DO UPDATE
SET canceller = EXCLUDED.canceller,
    cancelled_at = EXCLUDED.cancelled_at;

-- name: DeleteOrderByID :execresult
DELETE FROM orders WHERE id = $1;

-- name: UpdateOrderByID :execresult
UPDATE orders
SET id = $1,
    restaurant = $2,
    created_at = $3,
    customer = $4,
    customer_address = $5
WHERE id = $1;
