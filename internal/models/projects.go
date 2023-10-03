package models

import (
	"database/sql"
	"time"
)

type Projects struct {
	ID        string       `json:"id" db:"id" validate:"omitempty"`
	UserId    uint64       `json:"user_id" db:"user_id" validate:"required"`
	Name      string       `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time    `json:"created_at" db:"created_at" validate:"omitempty"`
	UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at" validate:"omitempty"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at" validate:"omitempty"`
}
