package auth

import (
	"context"
	"intro-ai/internal/models"
)

type Service interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	GetUserById(ctx context.Context, id uint64) (*models.User, error)
}
