-- +goose Up
-- +goose StatementBegin
ALTER TABLE images RENAME COLUMN filename TO name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE images RENAME COLUMN name TO filename;
-- +goose StatementEnd
