-- +goose Up
-- +goose StatementBegin
ALTER TABLE labels RENAME COLUMN project_id TO image_id;
ALTER TABLE labels DROP CONSTRAINT fk_projects_project_id;
ALTER TABLE labels ADD CONSTRAINT fk_images_image_id FOREIGN KEY (image_id) REFERENCES images (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE labels RENAME COLUMN image_id TO project_id;
ALTER TABLE labels DROP CONSTRAINT fk_images_image_id;
ALTER TABLE labels ADD CONSTRAINT fk_projects_project_id FOREIGN KEY (project_id) REFERENCES projects (id);
-- +goose StatementEnd
