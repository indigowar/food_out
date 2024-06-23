-- +goose Up
-- +goose StatementBegin
CREATE TABLE menu_dish(
	menu UUID NOT NULL,
	dish UUID NOT NULL,

	CONSTRAINT fk_menu
		FOREIGN KEY(menu)
		REFERENCES menus(id)
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
DROP TABLE menu_dish;
-- +goose StatementEnd
