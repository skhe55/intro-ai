package models

import (
	"time"
)

type Annotations struct {
	ID          string      `db:"id"`
	LabelID     string      `db:"label_id"`
	ImageID     string      `db:"image_id"`
	Coordinates [][]float64 `db:"coordinates"`
	CreatedAt   time.Time   `db:"created_at"`
}

type AnnotationDTO struct {
	ID          string      `json:"id" validate:"omitempty"`
	LabelID     string      `json:"label_id" validate:"required"`
	ImageID     string      `json:"image_id" validate:"required"`
	Coordinates [][]float64 `json:"coordinates" validate:"required"`
	CreatedAt   time.Time   `json:"created_at" validate:"omitempty"`
}

type AnnotationsWithLabelName struct {
	ID          string      `json:"id" validate:"omitempty"`
	LabelID     string      `json:"label_id" validate:"required"`
	LabelName   string      `json:"label_name" validate:"required"`
	ImageID     string      `json:"image_id" validate:"required"`
	Coordinates [][]float64 `json:"coordinates" validate:"required"`
	CreatedAt   time.Time   `json:"created_at" validate:"omitempty"`
}
