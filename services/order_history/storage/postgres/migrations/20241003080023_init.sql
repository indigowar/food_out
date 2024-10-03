-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders(
    id UUID NOT NULL,
    restaurant UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    customer UUID NOT NULL,
    customer_address VARCHAR(1024) NOT NULL,
    cooking_started_at TIMESTAMP,
    delivery_started_at TIMESTAMP,
    delivery_completed_at TIMESTAMP,

    PRIMARY KEY(id)
);

CREATE TABLE orders_acceptance(
    order_id UUID NOT NULL,
    manager UUID NOT NULL,
    accepted_at TIMESTAMP NOT NULL,

    PRIMARY KEY(order_id),

    FOREIGN KEY(order_id) REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE orders_couriers(
    order_id UUID NOT NULL,
    courier UUID NOT NULL,
    taken_at TIMESTAMP NOT NULL,

    PRIMARY KEY(order_id),

    FOREIGN KEY(order_id) REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE orders_payments(
    order_id UUID NOT NULL,
    transaction VARCHAR(1024) NOT NULL,
    payed_at TIMESTAMP NOT NULL,

    PRIMARY KEY(order_id),

    FOREIGN KEY(order_id) REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE orders_cancellations(
    order_id UUID NOT NULL,
    canceller UUID NOT NULL,
    cancelled_at TIMESTAMP NOT NULL,

    PRIMARY KEY(order_id),

    FOREIGN KEY(order_id) REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE products(
    id UUID NOT NULL,
    original_id UUID NOT NULL,
    restaurant UUID NOT NULL,
    order_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL CHECK(price > 0),
    description VARCHAR(1024) NOT NULL,
    
    PRIMARY KEY(id),

    FOREIGN KEY(order_id) REFERENCES orders(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
DROP TABLE orders_cancellations;
DROP TABLE orders_payments;
DROP TABLE orders_couriers;
DROP TABLE orders_acceptance;
DROP TABLE orders;
-- +goose StatementEnd

