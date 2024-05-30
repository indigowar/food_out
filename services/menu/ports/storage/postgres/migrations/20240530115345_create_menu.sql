-- +goose Up
-- +goose StatementBegin
CREATE TABLE menus(
	id UUID PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	image VARCHAR(1024) NOT NULL,
	restaurant UUID NOT NULL,
	FOREIGN KEY(restaurant) REFERENCES restaurants(id),

	is_deleted BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE menu_dish(
	menu UUID NOT NULL,
	dish UUID NOT NULL,

	FOREIGN KEY(menu) REFERENCES menus(id),
	FOREIGN KEY(dish) REFERENCES dishes(id),

	PRIMARY KEY(menu, dish)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menus;
DROP TABLE menu_dish;
-- +goose StatementEnd
