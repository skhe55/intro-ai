package server

import (
	authHttp "intro-ai/internal/auth/delivery/http"
	authRepository "intro-ai/internal/auth/repository"
	authService "intro-ai/internal/auth/service"
	imagesHttp "intro-ai/internal/images/delivery/http"
	imagesRepository "intro-ai/internal/images/repository"
	imagesService "intro-ai/internal/images/service"
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
	imagesRepository := imagesRepository.NewImagesRepository(s.db)

	authService := authService.NewAuthService(s.cfg, s.logger, authRepository)
	projectsService := projectsService.NewProjectsService(s.cfg, s.logger, projectsRepository)
	imagesService := imagesService.NewImagesService(s.cfg, s.logger, imagesRepository)

	myHttpError := httpError.NewHttpError()
	mw := middleware.NewMiddlewareManager(s.cfg, s.logger, myHttpError, authService)

	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.logger, myHttpError, authService)
	projectsHandlers := projectsHttp.NewProjectsHandlers(s.cfg, s.logger, myHttpError, projectsService)
	imagesHandlers := imagesHttp.NewImagesHandlers(s.cfg, s.logger, myHttpError, imagesService)

	authHttp.MapAuthRoutes("auth", authHandlers, mw, mux)
	projectsHttp.MapProjectsRoutes("projects", projectsHandlers, mw, mux)
	imagesHttp.MapImagesRoutes("images", imagesHandlers, mw, mux)
	return nil
}
