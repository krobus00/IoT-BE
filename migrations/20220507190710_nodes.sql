-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS nodes (
    id varchar(36) UNIQUE,
    city varchar(255),
    longitude decimal UNIQUE, 
    latitude decimal UNIQUE, 
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE nodes;
-- +goose StatementEnd