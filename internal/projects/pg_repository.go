package projects

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	GetAllProjects(ctx context.Context) ([]models.ProjectsWithImages, error)
	CreateProject(ctx context.Context, project *models.ProjectsDto, userId uint64) error
	DeleteProject(ctx context.Context, projectId string) error
}
