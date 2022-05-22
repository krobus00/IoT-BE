-- +goose Up
-- +goose StatementBegin
ALTER TABLE nodes
ADD unix_created_at BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP(),
ADD unix_updated_at BIGINT NOT NULL DEFAULT UNIX_TIMESTAMP(),
ADD unix_deleted_at BIGINT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sensors 
DROP COLUMN unix_created_at,
DROP COLUMN unix_updated_at,
DROP COLUMN unix_deleted_at;
-- +goose StatementEnd
