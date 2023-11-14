package models

import (
	"database/sql"
	"time"
)

type Labels struct {
	ID        string       `db:"id"`
	ImageId   string       `db:"image_id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type LabelDTO struct {
	ID        string    `json:"id" validate:"omitempty"`
	ImageId   string    `json:"image_id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"omitempty"`
}
