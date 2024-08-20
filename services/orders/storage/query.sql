-- name: InsertOrder :execresult
INSERT INTO orders (id, restaurant, customer, customer_address, created_at)
VALUES($1, $2, $3, $4, $5);

-- name: InsertOrderAcceptance :execresult
INSERT INTO order_acceptance(order_id, manager, accepted_at)
VALUES ($1, $2, $3);

-- name: InsertOrderPayment :execresult
INSERT INTO order_payment(order_id, transaction, payed_at)
VALUES($1, $2, $3);

-- name: InsertCancellation :execresult
INSERT INTO order_cancellation(order_id, canceller, cancelled_at)
VALUES ($1, $2, $3);

-- name: InsertProduct :execresult
INSERT INTO products(id, name, picture, price, description, order_id)
VALUES($1, $2, $3, $4, $5, $6);

-- name: InsertProductCategory :execresult
INSERT INTO product_categories(product, category)
VALUES ($1, $2);
