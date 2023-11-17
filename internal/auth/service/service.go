package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/auth"
	"intro-ai/internal/models"
	"intro-ai/pkg/logger"
	"intro-ai/pkg/utils"
)

type authService struct {
	cfg            *config.Config
	logger         logger.Logger
	authRepository auth.Repository
}

func NewAuthService(cfg *config.Config, logger logger.Logger, authRepository auth.Repository) auth.Service {
	return &authService{cfg: cfg, logger: logger, authRepository: authRepository}
}

func (u *authService) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	createdUser, err := u.authRepository.Register(ctx, user)

	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	token, expiresAt, err := utils.GenerateJWT(createdUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		UserName:  createdUser.UserName,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

func (u *authService) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	loggedUser, err := u.authRepository.Login(ctx, user)

	if err != nil {
		return nil, err
	}

	if err := loggedUser.ComparePasswords(user.Password); err != nil {
		return nil, err
	}

	token, expiresAt, err := utils.GenerateJWT(loggedUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		UserName:  loggedUser.UserName,
		Token:     token,
		ExpiresAt: expiresAt,
	}, nil
}

func (u *authService) GetUserById(ctx context.Context, id uint64) (*models.User, error) {
	user, err := u.authRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
