package labels

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	CreateLabel(ctx context.Context, label *models.LabelDTO) error
	DeleteLabel(ctx context.Context, labelId string) error
	GetLabelsByProjectId(ctx context.Context, projectId string) ([]models.LabelDTO, error)
}
