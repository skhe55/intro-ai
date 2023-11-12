-- +goose Up
-- +goose StatementBegin
ALTER TABLE images ADD CONSTRAINT constraint_uniq_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE images DROP CONSTRAINT IF EXISTS constraint_uniq_name;
-- +goose StatementEnd
