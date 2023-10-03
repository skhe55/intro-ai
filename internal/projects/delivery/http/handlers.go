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
	"net/http"
	"strings"
)

type projectsHandlers struct {
	cfg             *config.Config
	logger          logger.Logger
	projectsService projects.Service
}

func NewProjectsHandlers(
	cfg *config.Config,
	logger logger.Logger,
	projectsService projects.Service,
) projects.Handlers {
	return &projectsHandlers{
		cfg:             cfg,
		logger:          logger,
		projectsService: projectsService,
	}
}

func (h *projectsHandlers) GetAllProjects() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckHttpMethod(w, r, http.MethodGet); err != nil {
			h.logger.Error(err)
			return
		}

		myHttpError := httpError.NewHttpError(w)

		projects, err := h.projectsService.GetAllProjects(r.Context())
		if err != nil {
			myHttpError.InternalError()
			return
		}
		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, projects))
		if err != nil {
			h.logger.Error(err)
			myHttpError.InternalError()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *projectsHandlers) CreateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckHttpMethod(w, r, http.MethodPost); err != nil {
			h.logger.Error(err)
			return
		}

		myHttpError := httpError.NewHttpError(w)

		var project *models.Projects
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&project)
		if err != nil {
			myHttpError.NonInternalError(http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		claims, err := utils.ExtractJWTFromRequest(r)
		if err != nil {
			h.logger.Error(err)
			myHttpError.InternalError()
			return
		}

		if err := h.projectsService.CreateProject(r.Context(), project, uint64(claims["id"].(float64))); err != nil {
			myHttpError.InternalError()
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))
		if err != nil {
			h.logger.Error(err)
			myHttpError.InternalError()
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *projectsHandlers) DeleteProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckHttpMethod(w, r, http.MethodDelete); err != nil {
			h.logger.Error(err)
			return
		}

		myHttpError := httpError.NewHttpError(w)

		projectId := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

		if projectId == "" {
			myHttpError.NonInternalError(http.StatusNotFound, httpError.WRONG_ID)
			return
		}

		if err := h.projectsService.DeleteProject(r.Context(), projectId); err != nil {
			myHttpError.InternalError()
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))
		if err != nil {
			h.logger.Error(err)
			myHttpError.InternalError()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
