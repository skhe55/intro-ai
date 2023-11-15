package repository

import (
	"context"
	"encoding/json"
	"intro-ai/internal/annotations"
	"intro-ai/internal/models"
	"intro-ai/pkg/utils"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type annotationsRepository struct {
	db *sqlx.DB
}

func NewAnnotationsRepository(db *sqlx.DB) annotations.Repository {
	return &annotationsRepository{db: db}
}

func (r *annotationsRepository) CreateAnnotation(ctx context.Context, dto *models.AnnotationDTO) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	res, err := json.Marshal(dto.Coordinates)
	if err != nil {
		return err
	}

	coordinatesToString := string(res[:])

	coordinatesToString = strings.ReplaceAll(coordinatesToString, "[", "{")
	coordinatesToString = strings.ReplaceAll(coordinatesToString, "]", "}")

	if _, err := conn.ExecContext(
		ctx,
		"INSERT INTO annotations (label_id, image_id, coordinates, created_at) VALUES ($1, $2, $3, $4)",
		dto.LabelID,
		dto.ImageID,
		coordinatesToString,
		time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}

	return nil
}

func (r *annotationsRepository) DeleteAnnotation(ctx context.Context, labelId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE annotations SET deleted_at = $1::timestamp WHERE id = $2",
		time.Now().Format(time.RFC3339),
		labelId,
	); err != nil {
		return err
	}

	return nil
}

func (r *annotationsRepository) GetAnnotationsByImageId(ctx context.Context, imageId string) ([]models.AnnotationsWithLabelName, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var annotations []AnnotationsWithLabelNameDB

	if err := conn.SelectContext(
		ctx,
		&annotations,
		"SELECT a.id, a.label_id, l.name as label_name, a.coordinates, a.created_at FROM annotations as a INNER JOIN labels as l on a.label_id = l.id WHERE a.image_id = $1 and a.deleted_at iS NULL",
		imageId,
	); err != nil {
		return nil, err
	}

	return utils.Map[AnnotationsWithLabelNameDB, models.AnnotationsWithLabelName](annotations, func(item AnnotationsWithLabelNameDB, _ int) models.AnnotationsWithLabelName {
		return models.AnnotationsWithLabelName{
			ID:          item.ID,
			LabelID:     item.LabelID,
			LabelName:   item.LabelName,
			Coordinates: utils.ConvertStringToFloat64SliceOfSlices(string(item.Coordinates)),
			CreatedAt:   item.CreatedAt,
		}
	}), nil
}

func (r *annotationsRepository) GetAnnotationsByLabelId(ctx context.Context, labelId string) ([]models.Annotations, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var annotations []AnnotationsDB

	if err := conn.SelectContext(
		ctx,
		&annotations,
		"SELECT id, label_id, coordinates, created_at FROM annotations WHERE label_id = $1 and deleted_at IS NULL",
		labelId,
	); err != nil {
		return nil, err
	}

	return utils.Map[AnnotationsDB, models.Annotations](annotations, func(item AnnotationsDB, _ int) models.Annotations {
		return models.Annotations{
			ID:          item.ID,
			LabelID:     item.LabelID,
			Coordinates: utils.ConvertStringToFloat64SliceOfSlices(string(item.Coordinates)),
			CreatedAt:   item.CreatedAt,
		}
	}), nil
}

type AnnotationsDB struct {
	ID          string    `db:"id"`
	LabelID     string    `db:"label_id"`
	Coordinates []uint8   `db:"coordinates"`
	CreatedAt   time.Time `db:"created_at"`
}

type AnnotationsWithLabelNameDB struct {
	ID          string    `db:"id"`
	LabelID     string    `db:"label_id"`
	LabelName   string    `db:"label_name"`
	Coordinates []uint8   `db:"coordinates"`
	CreatedAt   time.Time `db:"created_at"`
}
