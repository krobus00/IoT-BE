-- +goose Up
-- +goose StatementBegin
ALTER TABLE nodes ADD model_url LONGTEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE nodes DROP COLUMN model_url;
-- +goose StatementEnd
