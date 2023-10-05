package images

import (
	"context"
	"intro-ai/internal/models"
)

type Service interface {
	GetAllImagesByProjectId(ctx context.Context, projectId string) ([]models.ImagesDTO, error)
	CreateImage(ctx context.Context, image *models.ImagesDTO) error
	DeleteImage(ctx context.Context, imageId string) error
}
