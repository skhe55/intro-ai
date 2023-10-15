package http

import (
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/models"
	"intro-ai/internal/projects"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"intro-ai/pkg/utils/httpHelper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type projectsHandlers struct {
	cfg             *config.Config
	logger          logger.Logger
	httpError       httpError.HttpError
	projectsService projects.Service
}

func NewProjectsHandlers(
	cfg *config.Config,
	logger logger.Logger,
	httpError httpError.HttpError,
	projectsService projects.Service,
) projects.Handlers {
	return &projectsHandlers{
		cfg:             cfg,
		logger:          logger,
		httpError:       httpError,
		projectsService: projectsService,
	}
}

func (h *projectsHandlers) GetAllProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects, err := h.projectsService.GetAllProjects(r.Context())
		if err != nil {
			h.httpError.InternalError(w)
			return
		}
		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, projects))
		if err != nil {
			h.logger.Error(err)
			h.httpError.InternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *projectsHandlers) CreateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var project *models.Projects
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&project)
		if err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(project); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		claims, err := utils.ExtractJWTFromRequest(r)
		if err != nil {
			h.logger.Error(err)
			h.httpError.InternalError(w)
			return
		}

		if err := h.projectsService.CreateProject(r.Context(), project, uint64(claims["id"].(float64))); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *projectsHandlers) DeleteProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)

		if projectId == "" {
			h.httpError.NonInternalError(w, http.StatusNotFound, httpError.WRONG_ID)
			return
		}

		if err := h.projectsService.DeleteProject(r.Context(), projectId); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
