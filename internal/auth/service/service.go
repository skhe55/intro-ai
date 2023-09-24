package service

import (
	"context"
	"intro-ai/config"
	"intro-ai/internal/auth"
	"intro-ai/internal/models"
	"intro-ai/pkg/utils"
)

type authService struct {
	cfg            *config.Config
	authRepository auth.Repository
}

func NewAuthService(cfg *config.Config, authRepository auth.Repository) auth.Service {
	return &authService{cfg: cfg, authRepository: authRepository}
}

func (u *authService) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	createdUser, err := u.authRepository.Register(ctx, user)

	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	token, err := utils.GenerateJWT(user, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		UserName: createdUser.UserName,
		Token:    token,
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

	token, err := utils.GenerateJWT(user, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		UserName: loggedUser.UserName,
		Token:    token,
	}, nil
}
