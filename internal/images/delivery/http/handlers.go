package http

import (
	"bytes"
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/images"
	"intro-ai/internal/models"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"intro-ai/pkg/utils/httpHelper"
	"io"
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

func (h *imagesHandlers) GetImageById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		imageId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)

		if imageId == "" {
			h.httpError.NonInternalError(w, http.StatusNotFound, httpError.WRONG_ID)
			return
		}
		image, err := h.imagesService.GetImageById(r.Context(), imageId)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(response.StatusOK, image))
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
			h.logger.Error(err)
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(image); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		imageId, err := h.imagesService.CreateImage(r.Context(), image)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, imageId))

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

		var dto *models.ImageDeleteDto
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&dto); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		validator := validator.New()
		if err := validator.Struct(dto); err != nil {
			h.httpError.NonInternalError(w, http.StatusBadRequest, err.Error())
			return
		}

		err := h.imagesService.DeleteImage(r.Context(), imageId, dto)
		if err != nil {
			h.httpError.InternalError(w)
			return
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func (h *imagesHandlers) UploadImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectIds, ok := r.URL.Query()["projectId"]
		imageId := httpHelper.RetriveIdFromUrlPath(r.URL.Path)
		if imageId == "" || !ok {
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_ID)
			return
		}

		readForm, err := r.MultipartReader()
		if err != nil {
			h.logger.Errorf("error occured: %v", err)
			h.httpError.NonInternalError(w, http.StatusBadRequest, httpError.WRONG_DTO)
			return
		}

		for {
			part, errPart := readForm.NextPart()
			if errPart == io.EOF {
				break
			}
			if part.FormName() == "file" {
				buf := new(bytes.Buffer)
				buf.ReadFrom(part)
				mimeType, err := utils.ConvertToValidMimeType(part.Header["Content-Type"][0])
				if err != nil {
					h.httpError.NonInternalError(w, http.StatusBadRequest, "Not recognized file extension. Supported extensions: jpg, png.")
					return
				}
				if err := h.imagesService.UploadImage(r.Context(), imageId, projectIds[0], buf, mimeType); err != nil {
					h.httpError.InternalError(w)
					return
				}
				break
			}
		}

		res, _ := utils.ToJSON[response.Response](response.OK(response.StatusOK, nil))

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
