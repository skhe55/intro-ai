package http

import (
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/images"
	"intro-ai/internal/models"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"intro-ai/pkg/utils/httpHelper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type imagesHandlers struct {
	cfg           *config.Config
	logger        logger.Logger
	httpError     httpError.HttpError
	imagesService images.Service
}

func NewImagesHandlers(
	cfg *config.Config,
	logger logger.Logger,
	httpError httpError.HttpError,
	imagesService images.Service,
) images.Handlers {
	return &imagesHandlers{
		cfg:           cfg,
		logger:        logger,
		httpError:     httpError,
		imagesService: imagesService,
	}
}

func (h *imagesHandlers) GetAllImagesByProjectId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)

		if projectId == "" {
			h.httpError.NonInternalError(w, http.StatusNotFound, httpError.WRONG_ID)
			return
		}

		images, err := h.imagesService.GetAllImagesByProjectId(r.Context(), projectId)

		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, images))
		if err != nil {
			h.logger.Error(err)
			h.httpError.InternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *imagesHandlers) CreateImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var image *models.ImagesDTO
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&image)
		if err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(image); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.imagesService.CreateImage(r.Context(), image); err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))
		if err != nil {
			h.logger.Error(err)
			h.httpError.InternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *imagesHandlers) DeleteImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)
		if imageId == "" {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_ID)
			return
		}

		err := h.imagesService.DeleteImage(r.Context(), imageId)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
