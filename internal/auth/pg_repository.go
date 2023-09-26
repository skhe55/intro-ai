package auth

import (
	"context"
	"intro-ai/internal/models"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, id uint64) (*models.User, error)
}
