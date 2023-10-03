package repository

import (
	"context"
	"intro-ai/internal/models"
	"intro-ai/internal/projects"
	"time"

	"github.com/jmoiron/sqlx"
)

type projectsRepository struct {
	db *sqlx.DB
}

func NewProjectsRepository(db *sqlx.DB) projects.Repository {
	return &projectsRepository{db: db}
}

func (r *projectsRepository) GetAllProjects(ctx context.Context) ([]models.Projects, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var projects []models.Projects

	if err := r.db.Select(
		&projects,
		"SELECT * FROM projects WHERE deleted_at IS NULL",
	); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectsRepository) CreateProject(ctx context.Context, project *models.Projects, userId uint64) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"INSERT INTO projects (user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		userId,
		project.Name,
		time.Now().UTC().Format(time.RFC3339),
		time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}

	return nil
}

func (r *projectsRepository) DeleteProject(ctx context.Context, projectId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE projects SET deleted_at = $1::timestamp WHERE id = $2",
		time.Now().UTC().Format(time.RFC3339),
		projectId,
	); err != nil {
		return err
	}

	return nil
}
