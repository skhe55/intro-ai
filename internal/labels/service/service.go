package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/labels"
	"intro-ai/internal/models"
	"intro-ai/pkg/logger"
)

type labelsService struct {
	cfg              *config.Config
	logger           logger.Logger
	labelsRepository labels.Repository
}

func NewLabelsService(
	cfg *config.Config,
	logger logger.Logger,
	labelsRepository labels.Repository,
) labels.Service {
	return &labelsService{
		cfg:              cfg,
		logger:           logger,
		labelsRepository: labelsRepository,
	}
}

func (s *labelsService) CreateLabel(ctx context.Context, labelDTO *models.LabelDTO) error {
	err := s.labelsRepository.CreateLabel(ctx, labelDTO)
	if err != nil {
		s.logger.Errorf("unable create label in db: %v", err)
		return err
	}

	return nil
}

func (s *labelsService) DeleteLabel(ctx context.Context, labelId string) error {
	err := s.labelsRepository.DeleteLabel(ctx, labelId)
	if err != nil {
		s.logger.Errorf("unable delete label from db: %v", err)
		return err
	}
	return nil
}
