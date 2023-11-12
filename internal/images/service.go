package images

import (
	"context"
	"intro-ai/internal/models"
	"io"
)

type Service interface {
	GetAllImagesByProjectId(ctx context.Context, projectId string) ([]models.ImagesDTO, error)
	CreateImage(ctx context.Context, image *models.ImagesDTO) (string, error)
	DeleteImage(ctx context.Context, imageId string, dto *models.ImageDeleteDto) error
	UploadImage(ctx context.Context, imageId string, projectId string, file io.Reader, mimeType string) error
	DeleteImagesByProjectId(ctx context.Context, projectId string) error
	GetImageById(ctx context.Context, imageId string) (*models.ImagesDTO, error)
}
