package middleware

import (
	"intro-ai/config"
	"intro-ai/internal/auth"
	"intro-ai/pkg/logger"
)

type MiddlewareManager struct {
	authService auth.Service
	cfg         *config.Config
	logger      logger.Logger
}

func NewMiddlewareManager(cfg *config.Config, logger logger.Logger, authService auth.Service) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, logger: logger, authService: authService}
}
