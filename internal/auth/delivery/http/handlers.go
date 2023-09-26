package http

import (
	"encoding/json"
	"intro-ai/config"
	"intro-ai/internal/auth"
	"intro-ai/internal/models"
	"intro-ai/internal/server/response"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
	"intro-ai/pkg/utils/httpError"
	"net/http"
	"strings"
)

var (
	COMMON_BAD_REQUEST_MESSAGE = "Проверьте значения переданные в дто."
	COMMON_SUCCESS_MESSAGE     = "Success"
)

type authHandlers struct {
	cfg         *config.Config
	logger      logger.Logger
	authService auth.Service
}

func NewAuthHandlers(cfg *config.Config, logger logger.Logger, authService auth.Service) auth.Handlers {
	return &authHandlers{cfg: cfg, logger: logger, authService: authService}
}

func (h *authHandlers) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckHttpMethod(w, r, http.MethodPost); err != nil {
			return
		}

		var user models.User

		httpError := httpError.NewHttpError(w)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&user)
		if err != nil {
			httpError.NonInternalError(http.StatusBadRequest, COMMON_BAD_REQUEST_MESSAGE)
			return
		}

		createdUser, err := h.authService.Register(r.Context(), &user)
		if err != nil {
			h.logger.Errorf("User already exist, err: %v", err)
			if strings.Contains(err.Error(), "пользователь уже существует") {
				httpError.NonInternalError(http.StatusBadRequest, "Пользователь уже существует.")
			} else {
				httpError.InternalError()
			}
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(COMMON_SUCCESS_MESSAGE, createdUser))
		if err != nil {
			httpError.InternalError()
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

func (h *authHandlers) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.CheckHttpMethod(w, r, http.MethodPost); err != nil {
			return
		}

		var user models.User

		httpError := httpError.NewHttpError(w)

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&user)
		if err != nil {
			httpError.NonInternalError(http.StatusBadRequest, COMMON_BAD_REQUEST_MESSAGE)
			return
		}

		loggedUser, err := h.authService.Login(r.Context(), &user)
		if err != nil {
			h.logger.Error(err)
			httpError.InternalError()
			return
		}

		res, err := utils.ToJSON[response.Response](response.OK(COMMON_SUCCESS_MESSAGE, loggedUser))
		if err != nil {
			httpError.InternalError()
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
