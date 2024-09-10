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
