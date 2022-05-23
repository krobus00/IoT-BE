-- +goose Up
-- +goose StatementBegin
ALTER TABLE nodes 
DROP COLUMN created_at,
CHANGE unix_created_at created_at BIGINT DEFAULT UNIX_TIMESTAMP() NOT NULL,
DROP COLUMN updated_at,
CHANGE unix_updated_at updated_at BIGINT DEFAULT UNIX_TIMESTAMP() NOT NULL,
DROP COLUMN deleted_at,
CHANGE unix_deleted_at deleted_at BIGINT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
