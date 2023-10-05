package repository

import (
	"context"
	"intro-ai/internal/images"
	"intro-ai/internal/models"
	"intro-ai/pkg/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

type imagesRepository struct {
	db *sqlx.DB
}

func NewImagesRepository(db *sqlx.DB) images.Repository {
	return &imagesRepository{db: db}
}

func (r *imagesRepository) GetAllImagesByProjectId(ctx context.Context, projectId string) ([]models.ImagesDTO, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var images []models.Images

	if err := conn.SelectContext(
		ctx,
		&images,
		"SELECT * FROM images WHERE deleted_at IS NULL and project_id = $1",
		projectId,
	); err != nil {
		return nil, err
	}

	return utils.Map[models.Images, models.ImagesDTO](images, func(item models.Images, _ int) models.ImagesDTO {
		return models.ImagesDTO{
			ID:          item.ID,
			ProjectId:   item.ProjectId,
			FileName:    item.FileName,
			PathToImage: item.PathToImage.String,
			Coordinates: item.Coordinates,
			CreatedAt:   item.CreatedAt,
		}
	}), nil
}

func (r *imagesRepository) CreateImage(ctx context.Context, image *models.ImagesDTO) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"INSERT INTO images (project_id, filename, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		image.ProjectId,
		image.FileName,
		time.Now().UTC().Format(time.RFC3339),
		time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}

	return nil
}

func (r *imagesRepository) DeleteImage(ctx context.Context, imageId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE images SET deleted_at = $1::timestamp WHERE id = $1",
		imageId,
	); err != nil {
		return err
	}
	return nil
}
