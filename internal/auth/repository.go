package auth

import (
	"context"
	"kienmatu/go-todos/models"
)

const CtxUserKey = "user"

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
