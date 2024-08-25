-- +goose Up
-- +goose StatementBegin
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

CREATE TABLE products_categories(
    product UUID NOT NULL,
    category VARCHAR(255) NOT NULL,
    
    PRIMARY KEY(product, category),
    
    FOREIGN KEY(product) REFERENCES products(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products_categories;
DROP TABLE products;
-- +goose StatementEnd
