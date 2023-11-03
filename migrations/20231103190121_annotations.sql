-- +goose Up
-- +goose StatementBegin
CREATE TABLE annotations (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    label_id UUID NOT NULL,
    coordinates FLOAT[][] NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT fk_labels_label_id
        FOREIGN KEY (label_id)
            REFERENCES labels (id)
            ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS annotations;
-- +goose StatementEnd
