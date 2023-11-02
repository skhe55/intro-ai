package http

import (
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/labels"
	"intro-ai/internal/models"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"intro-ai/pkg/utils/httpHelper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type labelsHandlers struct {
	cfg           *config.Config
	logger        logger.Logger
	httpError     httpError.HttpError
	labelsService labels.Service
}

func NewLabelsHandlers(
	cfg *config.Config,
	logger logger.Logger,
	httpError httpError.HttpError,
	labelsService labels.Service,
) labels.Handlers {
	return &labelsHandlers{
		cfg:           cfg,
		logger:        logger,
		httpError:     httpError,
		labelsService: labelsService,
	}
}

func (h *labelsHandlers) CreateLabel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var label *models.LabelDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&label)
		if err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(label); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.labelsService.CreateLabel(r.Context(), label); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *labelsHandlers) DeleteLabel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		labelId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)
		if labelId == "" {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_ID)
			return
		}

		err := h.labelsService.DeleteLabel(r.Context(), labelId)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
