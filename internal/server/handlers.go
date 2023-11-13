package server

import (
	annotationsHttp "intro-ai/internal/annotations/delivery/http"
	annotationsRepository "intro-ai/internal/annotations/repository"
	annotationsService "intro-ai/internal/annotations/service"
	authHttp "intro-ai/internal/auth/delivery/http"
	authRepository "intro-ai/internal/auth/repository"
	authService "intro-ai/internal/auth/service"
	imagesHttp "intro-ai/internal/images/delivery/http"
	imagesRepository "intro-ai/internal/images/repository"
	imagesService "intro-ai/internal/images/service"
	labelsHttp "intro-ai/internal/labels/delivery/http"
	labelsRepository "intro-ai/internal/labels/repository"
	labelsService "intro-ai/internal/labels/service"
	"intro-ai/internal/middleware"
	projectsHttp "intro-ai/internal/projects/delivery/http"
	projectsRepository "intro-ai/internal/projects/repository"
	projectsService "intro-ai/internal/projects/service"

	"intro-ai/pkg/utils/httpError"
	"net/http"
)

func (s *Server) MapHandlers(mux *http.ServeMux) error {
	fileServer := http.FileServer(http.Dir("./ftpdata/ikbkr/"))

	authRepository := authRepository.NewAuthRepository(s.db)
	projectsRepository := projectsRepository.NewProjectsRepository(s.db)
	imagesRepository := imagesRepository.NewImagesRepository(s.db)
	labelsRepository := labelsRepository.NewLabelsRepository(s.db)
	annotationsRepository := annotationsRepository.NewAnnotationsRepository(s.db)

	authService := authService.NewAuthService(s.cfg, s.logger, authRepository)
	projectsService := projectsService.NewProjectsService(s.cfg, s.logger, projectsRepository)
	imagesService := imagesService.NewImagesService(s.cfg, s.logger, imagesRepository)
	labelsService := labelsService.NewLabelsService(s.cfg, s.logger, labelsRepository)
	annotationsService := annotationsService.NewAnnotationsService(s.cfg, s.logger, annotationsRepository)

	myHttpError := httpError.NewHttpError()
	mw := middleware.NewMiddlewareManager(s.cfg, s.logger, myHttpError, authService)

	authHandlers := authHttp.NewAuthHandlers(s.cfg, s.logger, myHttpError, authService)
	projectsHandlers := projectsHttp.NewProjectsHandlers(s.cfg, s.logger, myHttpError, projectsService, imagesService)
	imagesHandlers := imagesHttp.NewImagesHandlers(s.cfg, s.logger, myHttpError, imagesService)
	labelsHandlers := labelsHttp.NewLabelsHandlers(s.cfg, s.logger, myHttpError, labelsService)
	annotationsHandlers := annotationsHttp.NewAnnotationsHandlers(s.cfg, s.logger, myHttpError, annotationsService)

	authHttp.MapAuthRoutes("auth", authHandlers, mw, mux)
	projectsHttp.MapProjectsRoutes("projects", projectsHandlers, mw, mux)
	imagesHttp.MapImagesRoutes("images", imagesHandlers, mw, mux)
	labelsHttp.MapLabelsRoutes("labels", labelsHandlers, mw, mux)
	annotationsHttp.MapAnnotationsRoutes("annotations", annotationsHandlers, mw, mux)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return nil
}
