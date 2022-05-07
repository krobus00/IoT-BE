-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sensors (
    id varchar(36) UNIQUE,
    node_id varchar(255),
    humidity float,
    temperature float,
    heat_index float,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sensors;
-- +goose StatementEnd