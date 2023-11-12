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
			Name:        item.Name,
			PathToImage: item.PathToImage.String,
			CreatedAt:   item.CreatedAt,
		}
	}), nil
}

func (r *imagesRepository) CreateImage(ctx context.Context, imageDto *models.ImagesDTO) (string, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	var imageId string

	row := conn.QueryRowxContext(
		ctx,
		"INSERT INTO images (project_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) returning id",
		imageDto.ProjectId,
		imageDto.Name,
		time.Now().UTC().Format(time.RFC3339),
		time.Now().UTC().Format(time.RFC3339),
	)
	if row.Err() != nil {
		return "", err
	}
	if err := row.Scan(&imageId); err != nil {
		return "", err
	}
	return imageId, nil
}

func (r *imagesRepository) GetImageById(ctx context.Context, imageId string) (*models.ImagesDTO, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var image models.Images

	row := conn.QueryRowxContext(
		ctx,
		"SELECT * FROM images WHERE id = $1",
		imageId,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	if err := row.StructScan(&image); err != nil {
		return nil, err
	}

	return &models.ImagesDTO{
		ID:          image.ID,
		Name:        image.Name,
		ProjectId:   image.ProjectId,
		PathToImage: image.PathToImage.String,
		CreatedAt:   image.CreatedAt,
	}, nil
}

func (r *imagesRepository) UploadImage(ctx context.Context, imageId string, pathToImage string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE images SET path_to_image = $1 WHERE id = $2",
		pathToImage,
		imageId,
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
		"UPDATE images SET deleted_at = $1::timestamp, path_to_image = '' WHERE id = $2",
		time.Now().Format(time.RFC3339),
		imageId,
	); err != nil {
		return err
	}
	return nil
}

func (r *imagesRepository) DeleteImagesByProjectId(ctx context.Context, projectId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"DELETE FROM images WHERE project_id = $1",
		projectId,
	); err != nil {
		return err
	}

	return nil
}
