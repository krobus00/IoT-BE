-- +goose Up
-- +goose StatementBegin
ALTER TABLE nodes ADD fallback_model_url LONGTEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE nodes DROP COLUMN fallback_model_url;
-- +goose StatementEnd
