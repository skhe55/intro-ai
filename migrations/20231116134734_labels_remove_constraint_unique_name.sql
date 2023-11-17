-- +goose Up
-- +goose StatementBegin
ALTER TABLE labels DROP CONSTRAINT IF EXISTS labels_name_key;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE labels ADD CONSTRAINT labels_name_key UNIQUE (name);
-- +goose StatementEnd
