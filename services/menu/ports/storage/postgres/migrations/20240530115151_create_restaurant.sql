-- +goose Up
-- +goose StatementBegin
CREATE TABLE restaurants(
	id UUID PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE restaurants;
-- +goose StatementEnd
