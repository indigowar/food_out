CREATE TABLE orders(
    id UUID PRIMARY KEY,

    restaurant UUID NOT NULL,
    customer UUID NOT NULL,
    customer_address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE order_acceptance(
    order_id UUID PRIMARY KEY,
    FOREIGN KEY(order_id) REFERENCES orders(id),

    manager UUID NOT NULL,
    accepted_at TIMESTMAP NOT NULL
);

CREATE TABLE order_courier(
    order_id UUID PRIMARY KEY,
    FOREIGN KEY(order_id) REFERENCES orders(id),

    courier UUID NOT NULL,
    taken_at TIMESTAMP NOT NULL
);

CREATE TABLE order_payment(
    order_id UUID PRIMARY KEY,
    FOREIGN KEY(order_id) REFERENCES orders(id),

    transaction VARCHAR(255) NOT NULL,
    payed_at TIMESTAMP NOT NULL
);

CREATE TABLE order_cancellation(
    order_id UUID PRIMARY KEY,
    FOREIGN KEY(order_id) REFERENCES orders(id),

    canceller UUID NOT NULL,
    cancelled_at TIMESTAMP NOT NULL
);

CREATE TABLE products(
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL CHECK(price > 0),
    description VARCHAR(255) NOT NULL,

    order_id UUID REFERENCES orders(id) NOT NULL
);

CREATE TABLE product_categories(
    product UUID REFERENCES products(id) NOT NULL,
    category VARCHAR(255) NOT NULL,

    UNIQUE(product, category)
);
