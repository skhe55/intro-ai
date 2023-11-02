-- +goose Up
-- +goose StatementBegin
CREATE TABLE labels (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    project_id UUID NOT NULL,
    name varchar(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_projects_project_id
        FOREIGN KEY(project_id)
            REFERENCES projects (id)
            ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS labels;
-- +goose StatementEnd
