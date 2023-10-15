package images

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	GetAllImagesByProjectId(ctx context.Context, projectId string) ([]models.ImagesDTO, error)
	CreateImage(ctx context.Context, image *models.ImagesDTO) error
	UploadImage(ctx context.Context, imageId string, pathToImage string) error
	DeleteImage(ctx context.Context, imageId string) error
}
