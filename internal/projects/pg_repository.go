package projects

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	GetAllProjects(ctx context.Context) ([]models.Projects, error)
	CreateProject(ctx context.Context, project *models.Projects, userId uint64) error
	DeleteProject(ctx context.Context, projectId string) error
}
