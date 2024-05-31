-- +goose Up
-- +goose StatementBegin
CREATE TABLE menus(
	id UUID PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	image VARCHAR(1024) NOT NULL,
	restaurant UUID NOT NULL,
	FOREIGN KEY(restaurant) REFERENCES restaurants(id)
);

CREATE TABLE menu_dish(
	menu UUID NOT NULL,
	dish UUID NOT NULL,

	CONSTRAINT fk_menu
		FOREIGN KEY(menu)
		REFERENCES menu(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE,

	CONSTRAINT fk_dish
		FOREIGN KEY(dish)
		REFERENCES dishes(id)
		ON UPDATE CASCADE
		ON DELETE CASCADE,

	PRIMARY KEY(menu, dish)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE menus;
DROP TABLE menu_dish;
-- +goose StatementEnd
