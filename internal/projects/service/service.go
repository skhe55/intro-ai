package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/models"
	"intro-ai/internal/projects"
	"intro-ai/pkg/logger"
)

type projectsService struct {
	cfg                *config.Config
	logger             logger.Logger
	projectsRepository projects.Repository
}

func NewProjectsService(
	cfg *config.Config,
	logger logger.Logger,
	projectsRepository projects.Repository,
) projects.Service {
	return &projectsService{
		cfg:                cfg,
		logger:             logger,
		projectsRepository: projectsRepository,
	}
}

func (s *projectsService) GetAllProjects(ctx context.Context) ([]models.Projects, error) {
	projects, err := s.projectsRepository.GetAllProjects(ctx)

	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return projects, nil
}

func (s *projectsService) CreateProject(ctx context.Context, project *models.Projects, userId uint64) error {
	err := s.projectsRepository.CreateProject(ctx, project, userId)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}

func (s *projectsService) DeleteProject(ctx context.Context, projectId string) error {
	err := s.projectsRepository.DeleteProject(ctx, projectId)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	return nil
}
