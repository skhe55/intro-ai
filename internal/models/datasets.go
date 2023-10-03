package models

import (
	"database/sql"
	"time"
)

type Datasets struct {
	ID          string       `json:"id" db:"id" validate:"omitempty"`
	ProjectId   string       `json:"project_id" db:"project_id" validate:"required"`
	FileName    string       `json:"filename" db:"filename" validate:"required"`
	Coordinates [][]float32  `json:"coordinates" db:"coordinates" validate:"required"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at" validate:"omitempty"`
	UpdatedAt   sql.NullTime `json:"updated_at" db:"updated_at" validate:"omitempty"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at" validate:"omitempty"`
}
