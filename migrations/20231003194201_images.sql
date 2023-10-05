-- +goose Up
-- +goose StatementBegin
CREATE TABLE images (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    project_id UUID NOT NULL,
    filename VARCHAR(512),
    path_to_image VARCHAR(512),
    coordinates FLOAT[][],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_projects_project_id
        FOREIGN KEY(project_id)
            REFERENCES projects (id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS images;
-- +goose StatementEnd
