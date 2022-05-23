-- +goose Up
-- +goose StatementBegin
UPDATE nodes SET
unix_created_at = UNIX_TIMESTAMP(created_at), 
unix_updated_at  = UNIX_TIMESTAMP(updated_at),
unix_deleted_at = UNIX_TIMESTAMP(deleted_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
