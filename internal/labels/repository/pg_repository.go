package repository

import (
	"context"
	"intro-ai/internal/labels"
	"intro-ai/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type labelsRepository struct {
	db *sqlx.DB
}

func NewLabelsRepository(db *sqlx.DB) labels.Repository {
	return &labelsRepository{db: db}
}

func (r *labelsRepository) CreateLabel(ctx context.Context, labelDTO *models.LabelDTO) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"INSERT INTO labels (project_id, name, created_at) VALUES ($1, $2, $3)",
		labelDTO.ProjectId,
		labelDTO.Name,
		time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}

	return nil
}

func (r *labelsRepository) DeleteLabel(ctx context.Context, labelId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE labels SET deleted_at = $1::timestamp WHERE id = $2",
		time.Now().Format(time.RFC3339),
		labelId,
	); err != nil {
		return err
	}

	return nil
}
