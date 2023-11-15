package annotations

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	GetAnnotationsByLabelId(ctx context.Context, labelId string) ([]models.Annotations, error)
	GetAnnotationsByImageId(ctx context.Context, imageId string) ([]models.AnnotationsWithLabelName, error)
	CreateAnnotation(ctx context.Context, dto *models.AnnotationDTO) error
	DeleteAnnotation(ctx context.Context, annotationId string) error
}
