-- +goose Up
-- +goose StatementBegin
ALTER TABLE annotations ADD COLUMN image_id UUID;
ALTER TABLE annotations ADD CONSTRAINT fk_images_image_id FOREIGN KEY (image_id) REFERENCES images (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE annotations DROP CONSTRAINT fk_images_image_id;
ALTER TABLE annotations DROP COLUMN image_id;
-- +goose StatementEnd
