package labels

import (
	"context"
	"intro-ai/internal/models"
)

type Service interface {
	CreateLabel(ctx context.Context, labelDTO *models.LabelDTO) error
	DeleteLabel(ctx context.Context, labelId string) error
	GetLabelsByProjectId(ctx context.Context, projectId string) ([]models.LabelDTO, error)
}
