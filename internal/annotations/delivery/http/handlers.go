package http

import (
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/annotations"
	"intro-ai/internal/models"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"intro-ai/pkg/utils/httpHelper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type annotationsHandlers struct {
	cfg                *config.Config
	logger             logger.Logger
	httpError          httpError.HttpError
	annotationsService annotations.Service
}

func NewAnnotationsHandlers(
	cfg *config.Config,
	logger logger.Logger,
	httpError httpError.HttpError,
	annotationsService annotations.Service,
) annotations.Handlers {
	return &annotationsHandlers{
		cfg:                cfg,
		logger:             logger,
		httpError:          httpError,
		annotationsService: annotationsService,
	}
}

func (h *annotationsHandlers) CreateAnnotation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var annotation *models.AnnotationDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&annotation)
		if err != nil {
			h.logger.Error(err)
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(annotation); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.annotationsService.CreateAnnotation(r.Context(), annotation); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *annotationsHandlers) DeleteAnnotation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labelId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)
		if labelId == "" {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_ID)
			return
		}

		if err := h.annotationsService.DeleteAnnotation(r.Context(), labelId); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *annotationsHandlers) GetAnnotationsByLabelId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labelId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)
		if labelId == "" {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_ID)
			return
		}

		annotations, err := h.annotationsService.GetAnnotationsByLabelId(r.Context(), labelId)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, annotations))
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
