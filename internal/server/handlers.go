package server

import (
	authHttp "intro-ai/internal/auth/delivery/http"
	authRepository "intro-ai/internal/auth/repository"
	authService "intro-ai/internal/auth/service"
	"intro-ai/internal/middleware"
	projectsHttp "intro-ai/internal/projects/delivery/http"
	projectsRepository "intro-ai/internal/projects/repository"
	projectsService "intro-ai/internal/projects/service"
	"intro-ai/pkg/utils/httpError"
	"net/http"
)

func (s *Server) MapHandlers(mux *http.ServeMux) error {
	authRepository := authRepository.NewAuthRepository(s.db)
	projectsRepository := projectsRepository.NewProjectsRepository(s.db)

	authService := authService.NewAuthService(s.cfg, s.logger, authRepository)
	projectsService := projectsService.NewProjectsService(s.cfg, s.logger, projectsRepository)

	myHttpError := httpError.NewHttpError()
	mw := middleware.NewMiddlewareManager(s.cfg, s.logger, myHttpError, authService)

	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.logger, myHttpError, authService)
	projectsHandlers := projectsHttp.NewProjectsHandlers(s.cfg, s.logger, myHttpError, projectsService)

	authHttp.MapAuthRoutes("auth", authHandlers, mw, mux)
	projectsHttp.MapProjectsRoutes("projects", projectsHandlers, mw, mux)
	return nil
}
