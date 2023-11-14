package labels

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	CreateLabel(ctx context.Context, label *models.LabelDTO) error
	DeleteLabel(ctx context.Context, labelId string) error
	GetLabelsByImageId(ctx context.Context, imageId string) ([]models.LabelDTO, error)
}
