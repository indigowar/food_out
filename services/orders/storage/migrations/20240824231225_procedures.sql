-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION set_order_acceptance(
    in_order UUID,
    in_manager UUID,
    in_timestamp TIMESTAMP
) AS $$
BEGIN
    IF EXISTS( SELECT 1 FROM orders_acceptance WHERE order_id = in_order)
    BEGIN
        UPDATE orders_acceptance
        WHERE order_id = in_order
        SET manager = in_manager, accepted_at = in_timestamp;
    END
    ELSE
    BEGIN
        INSERT INTO orders_acceptance (order_id, manager, manager, accepted_at)
        VALUES(in_order, in_manager, in_timestamp);
    END
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION set_order_couriers(
    in_order UUID,
    in_courier UUID,
    in_timestamp TIMESTAMP
) AS $$ 
BEGIN
    IF EXISTS(SELECT 1 FROM orders_couriers WHERE order_id = in_order)
    BEGIN
        UPDATE orders_couriers
        WHERE order_id = in_order
        SET courier = in_courier, taken_at = in_timestamp;
    END
    ELSE BEGIN
        INSERT INTO orders_couriers(order_id, courier, taken_at)
        VALUES(in_order, in_courier, in_timestamp);
    END
END;
$$ LANGUAGE plpgsql;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION set_order_acceptance;
-- +goose StatementEnd
