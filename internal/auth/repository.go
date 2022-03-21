package auth

import (
	"context"
	"kienmatu/go-todos/internal/models"
)

const CtxUserKey = "userId"

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
