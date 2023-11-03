package models

import (
	"database/sql"
	"time"
)

type Images struct {
	ID          string         `json:"id" db:"id"`
	ProjectId   string         `json:"projectId" db:"project_id"`
	FileName    string         `json:"filename" db:"filename"`
	PathToImage sql.NullString `json:"pathToImage" db:"path_to_image"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}

type ImagesDTO struct {
	ID          string    `json:"id" validate:"omitempty" db:"id"`
	ProjectId   string    `json:"projectId" validate:"required" db:"project_id"`
	FileName    string    `json:"filename" validate:"required" db:"filename"`
	PathToImage string    `json:"pathToImage" validate:"omitempty" db:"path_to_image"`
	CreatedAt   time.Time `json:"created_at" validate:"omitempty" db:"created_at"`
}

type ImageDeleteDto struct {
	PathToImage string `json:"pathToImage" validate:"required"`
}
