package annotations

import (
	"context"
	"intro-ai/internal/models"
)

type Service interface {
	GetAnnotationsByLabelId(ctx context.Context, labelId string) ([]models.Annotations, error)
	CreateAnnotation(ctx context.Context, dto *models.AnnotationDTO) error
	DeleteAnnotation(ctx context.Context, labelId string) error
}
