-- Add migration script here

CREATE TABLE IF NOT EXISTS orders(
    id UUID PRIMARY KEY,
    customer UUID NOT NULL,
    customer_address VARCHAR(1024) NOT NULL,
    restaurant UUID NOT NULL,
    created_at TIMESTMAP NOT NULL,

    cooking_started_at TIMESTAMP,
    delivery_started_at TIMESTAMP,
    delivery_completed_at TIMESTAMP,

    UNIQUE(customer, restaurant, created_at)
);

CREATE TABLE IF NOT EXISTS order_acceptances(
    order UUID PRIMARY KEY,
    FOREIGN KEY(order) REFERENCES orders(id),

    manager UUID NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_couriers(
    order UUID PRIMARY KEY,
    FOREIGN KEY(order) REFERENCES orders(id),

    courier UUID NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_payments(
    order UUID PRIMARY KEY,
    FOREIGN KEY(order) REFERENCES orders(id),

    transaction VARCHAR(1024) NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS order_cancellations(
    order UUID PRIMARY KEY,
    FOREIGN KEY(order) REFERENCES orders(id),

    canceller UUID NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS products(
    id UUID PRIMARY KEY,
    order UUID REFERENCES orders(id) NOT NULL,
    restaurant UUID REFERENCES orders(restaurant) NOT NULL,
    name VARCHAR(256) NOT NULL,
    picture VARCHAR(1024) NOT NULL,
    price FLOAT NOT NULL CHECK (price > 0)
);

CREATE TABLE IF NOT EXISTS categories(
    product UUID REFERENCES products(id) NOT NULL,
    category VARCHAR(256) NOT NULL,
    UNIQUE(product, category)
);
