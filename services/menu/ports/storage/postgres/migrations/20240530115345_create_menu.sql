-- +goose Up
-- +goose StatementBegin
CREATE TABLE menus(
	id UUID PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	image VARCHAR(1024) NOT NULL,
	restaurant UUID NOT NULL,
	FOREIGN KEY(restaurant) REFERENCES restaurants(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menus;
-- +goose StatementEnd
