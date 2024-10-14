-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    phone VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(1024) NOT NULL,
    name VARCHAR(64),
    profile VARCHAR(1024)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
-- +goose StatementEnd
