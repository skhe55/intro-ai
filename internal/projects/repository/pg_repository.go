package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"intro-ai/internal/models"
	"intro-ai/internal/projects"
	"intro-ai/pkg/utils"
	"time"

	"github.com/jmoiron/sqlx"
)

type projectsRepository struct {
	db *sqlx.DB
}

func NewProjectsRepository(db *sqlx.DB) projects.Repository {
	return &projectsRepository{db: db}
}

func (r *projectsRepository) GetAllProjects(ctx context.Context) ([]models.ProjectsWithImages, error) {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var projects []ProjectsWithImagesDB

	if err := r.db.Select(
		&projects,
		"select projects.id, projects.name, JSON_AGG(images) as images from projects "+
			"left join images on projects.id = images.project_id and images.deleted_at is null "+
			"group by projects.id "+
			"having projects.deleted_at is null",
	); err != nil {
		return nil, err
	}

	return utils.Map[ProjectsWithImagesDB, models.ProjectsWithImages](projects, func(item ProjectsWithImagesDB, _ int) models.ProjectsWithImages {
		var dest interface{}
		var images []models.ImagesDTO

		json.Unmarshal(item.Images, &dest)
		switch v := dest.(type) {
		case []interface{}:
			if v[0] != nil {
				images = utils.Map[interface{}, models.ImagesDTO](v, func(mapItem interface{}, _ int) models.ImagesDTO {
					var pathToImage string
					var coordinates *[][]sql.NullFloat64
					var createdAt, _ = time.Parse(time.RFC3339, mapItem.(map[string]interface{})["created_at"].(string))

					if mapItem.(map[string]interface{})["path_to_image"] != nil {
						pathToImage = mapItem.(map[string]interface{})["path_to_image"].(string)
					}
					if mapItem.(map[string]interface{})["coordinates"] != nil {
						coordinates = mapItem.(map[string]interface{})["coordinates"].(*[][]sql.NullFloat64)
					}

					return models.ImagesDTO{
						ID:          mapItem.(map[string]interface{})["id"].(string),
						ProjectId:   mapItem.(map[string]interface{})["project_id"].(string),
						FileName:    mapItem.(map[string]interface{})["filename"].(string),
						PathToImage: pathToImage,
						Coordinates: coordinates,
						CreatedAt:   createdAt}
				})
			}
		}

		return models.ProjectsWithImages{
			ID:     item.ID,
			Name:   item.Name,
			Images: images,
		}
	}), nil
}

func (r *projectsRepository) CreateProject(ctx context.Context, project *models.ProjectsDto, userId uint64) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"INSERT INTO projects (user_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)",
		userId,
		project.Name,
		time.Now().UTC().Format(time.RFC3339),
		time.Now().UTC().Format(time.RFC3339),
	); err != nil {
		return err
	}

	return nil
}

func (r *projectsRepository) DeleteProject(ctx context.Context, projectId string) error {
	conn, err := r.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		"UPDATE projects SET deleted_at = $1::timestamp WHERE id = $2",
		time.Now().UTC().Format(time.RFC3339),
		projectId,
	); err != nil {
		return err
	}

	return nil
}

type ProjectsWithImagesDB struct {
	ID     string  `db:"id"`
	Name   string  `db:"name"`
	Images []uint8 `db:"images"`
}
