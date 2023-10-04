package middleware

import (
	"intro-ai/config"
	"intro-ai/internal/auth"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils/httpError"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MiddlewareManager struct {
	authService auth.Service
	cfg         *config.Config
	httpError   httpError.HttpError
	logger      logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, logger logger.Logger, httpError httpError.HttpError, authService auth.Service) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, logger: logger, httpError: httpError, authService: authService}
}

func (mw MiddlewareManager) Method(method string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				mw.httpError.NonInternalError(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
				return
			}

			f(w, r)
		}
	}
}

func (mw MiddlewareManager) Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
