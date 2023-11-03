-- +goose Up
-- +goose StatementBegin
ALTER TABLE images DROP COLUMN coordinates;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE images ADD COLUMN coordinates float[][];
-- +goose StatementEnd
