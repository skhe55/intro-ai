package projects

import (
	"context"
	"intro-ai/internal/models"
)

type Service interface {
	GetAllProjects(ctx context.Context, userId uint64) ([]models.ProjectsWithImages, error)
	CreateProject(ctx context.Context, project *models.ProjectsDto, userId uint64) error
	DeleteProject(ctx context.Context, projectId string) error
}
