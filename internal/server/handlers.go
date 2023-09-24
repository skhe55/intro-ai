package server

import (
	authHttp "intro-ai/internal/auth/delivery/http"
	"intro-ai/internal/auth/repository"
	"intro-ai/internal/auth/service"
)

func (s *Server) MapHandlers() error {
	authRepository := repository.NewAuthRepository(s.db)
	authService := service.NewAuthService(s.cfg, authRepository)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authService)
	authHttp.MapAuthRoutes("auth", authHandlers)
	return nil
}
