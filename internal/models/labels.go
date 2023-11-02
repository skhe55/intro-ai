package models

import (
	"database/sql"
	"time"
)

type Labels struct {
	ID        string       `db:"id"`
	ProjectId string       `db:"project_id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type LabelDTO struct {
	ID        string    `json:"id" validate:"omitempty"`
	ProjectId string    `json:"project_id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"omitempty"`
}
