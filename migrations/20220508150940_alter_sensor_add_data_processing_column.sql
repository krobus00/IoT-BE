-- +goose Up
-- +goose StatementBegin
ALTER TABLE sensors
    ADD COLUMN filter_humidity float AFTER humidity,
    ADD COLUMN filter_temperature float AFTER temperature,
    ADD COLUMN filter_heat_index float AFTER heat_index;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sensors
    DROP COLUMN filter_humidity float AFTER humidity,
    DROP COLUMN filter_temperature float AFTER temperature,
    DROP COLUMN filter_heat_index float AFTER heat_index;
-- +goose StatementEnd
