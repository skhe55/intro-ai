package service

import (
	"context"
	"fmt"
	"intro-ai/config"
	"intro-ai/internal/images"
	"intro-ai/internal/models"
	"intro-ai/pkg/logger"
	"io"

	"github.com/google/uuid"
	"github.com/jlaffaye/ftp"
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
		s.logger.Errorf("unable get images from db: %v", err)
		return nil, err
	}

	return images, nil
}

func (s *imagesService) CreateImage(ctx context.Context, image *models.ImagesDTO) error {
	if err := s.imagesRepository.CreateImage(ctx, image); err != nil {
		s.logger.Errorf("unable save image in db: %v", err)
		return err
	}
	return nil
}

func (s *imagesService) DeleteImage(ctx context.Context, imageId string, dto *models.ImageDeleteDto) error {
	err := s.imagesRepository.DeleteImage(ctx, imageId)
	if err != nil {
		s.logger.Errorf("unable delete image from db: %v", err)
		return err
	}

	conn, err := ftp.Dial(s.cfg.FtpConnectionString)
	if err != nil {
		s.logger.Errorf("unable connect to ftp server: %v", err)
		return err
	}
	defer conn.Quit()

	if err := conn.Login(s.cfg.FtpUserLogin, s.cfg.FtpUserPassword); err != nil {
		s.logger.Errorf("unable sign in to ftp server: %v", err)
		return err
	}

	if err := conn.Delete(dto.PathToImage); err != nil {
		s.logger.Errorf("failed delete image from ftp server: %v", err)
		return err
	}

	return nil
}

func (s *imagesService) UploadImage(ctx context.Context, imageId string, file io.Reader, mimeType string) error {
	c, err := ftp.Dial(s.cfg.FtpConnectionString)
	if err != nil {
		s.logger.Errorf("unable connect to ftp server: %v", err)
		return err
	}
	defer c.Quit()

	if err := c.Login(s.cfg.FtpUserLogin, s.cfg.FtpUserPassword); err != nil {
		s.logger.Errorf("unable sign in to ftp server: %v", err)
		return err
	}

	fileName := fmt.Sprintf("%v.%s", uuid.New(), mimeType)
	if err := c.Stor(fileName, file); err != nil {
		s.logger.Errorf("unable store to ftp server: %v", err)
		return err
	}

	if err := s.imagesRepository.UploadImage(ctx, imageId, fileName); err != nil {
		s.logger.Errorf("unable save image in db: %v", err)
		return err
	}

	return nil
}
