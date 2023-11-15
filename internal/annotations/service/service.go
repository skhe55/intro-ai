package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/annotations"
	"intro-ai/internal/models"
	"intro-ai/pkg/logger"
)

type annotationsService struct {
	cfg                   *config.Config
	logger                logger.Logger
	annotationsRepository annotations.Repository
}

func NewAnnotationsService(
	cfg *config.Config,
	logger logger.Logger,
	annotationsRepository annotations.Repository,
) annotations.Service {
	return &annotationsService{
		cfg:                   cfg,
		logger:                logger,
		annotationsRepository: annotationsRepository,
	}
}

func (s *annotationsService) CreateAnnotation(ctx context.Context, dto *models.AnnotationDTO) error {
	if err := s.annotationsRepository.CreateAnnotation(ctx, dto); err != nil {
		s.logger.Errorf("unable to create annotation: %v", err)
		return err
	}

	return nil
}

func (s *annotationsService) DeleteAnnotation(ctx context.Context, annotationId string) error {
	if err := s.annotationsRepository.DeleteAnnotation(ctx, annotationId); err != nil {
		s.logger.Errorf("unable to delete annotation: %v", err)
		return err
	}

	return nil
}

func (s *annotationsService) GetAnnotationsByLabelId(ctx context.Context, labelId string) ([]models.Annotations, error) {
	annotations, err := s.annotationsRepository.GetAnnotationsByLabelId(ctx, labelId)
	if err != nil {
		s.logger.Errorf("unable to get list of annotations: %v", err)
		return nil, err
	}

	return annotations, nil
}

func (s *annotationsService) GetAnnotationsByImageId(ctx context.Context, imageId string) ([]models.AnnotationsWithLabelName, error) {
	annotations, err := s.annotationsRepository.GetAnnotationsByImageId(ctx, imageId)
	if err != nil {
		s.logger.Errorf("unable to get list of annotations by image id: %v", err)
		return nil, err
	}

	return annotations, nil
}
