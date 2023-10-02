package server

import (
	authHttp "intro-ai/internal/auth/delivery/http"
	"intro-ai/internal/auth/repository"
	"intro-ai/internal/auth/service"
	"intro-ai/internal/middleware"
	"net/http"
)

func (s *Server) MapHandlers(mux *http.ServeMux) error {
	authRepository := repository.NewAuthRepository(s.db)

	authService := service.NewAuthService(s.cfg, s.logger, authRepository)

	mw := middleware.NewMiddlewareManager(s.cfg, s.logger, authService)

	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.logger, authService)

	authHttp.MapAuthRoutes("auth", authHandlers, mw, mux)
	return nil
}
