package server

import (
	authHttp "intro-ai/internal/auth/delivery/http"
)

func (s *Server) MapHandlers() error {
	authHttp.MapAuthRoutes("auth")
	return nil
}
