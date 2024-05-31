-- +goose Up
-- +goose StatementBegin
CREATE TABLE dishes(
	id UUID PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	image VARCHAR(1024) NOT NULL,
	price FLOAT NOT NULL CHECK ( price > 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dishes;
-- +goose StatementEnd
