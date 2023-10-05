package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/images"
	"intro-ai/internal/models"
	"intro-ai/pkg/logger"
)

type imagesService struct {
	cfg              *config.Config
	logger           logger.Logger
	imagesRepository images.Repository
}

func NewImagesService(
	cfg *config.Config,
	logger logger.Logger,
	imagesRepository images.Repository,
) images.Service {
	return &imagesService{
		cfg:              cfg,
		logger:           logger,
		imagesRepository: imagesRepository,
	}
}

func (s *imagesService) GetAllImagesByProjectId(ctx context.Context, projectId string) ([]models.ImagesDTO, error) {
	images, err := s.imagesRepository.GetAllImagesByProjectId(ctx, projectId)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}

	return images, nil
}

func (s *imagesService) CreateImage(ctx context.Context, image *models.ImagesDTO) error {
	err := s.imagesRepository.CreateImage(ctx, image)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}

func (s *imagesService) DeleteImage(ctx context.Context, imageId string) error {
	err := s.imagesRepository.DeleteImage(ctx, imageId)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}
