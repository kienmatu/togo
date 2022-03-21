package auth

import (
	"context"
	"kienmatu/go-todos/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, limit int) (*models.User, error)
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
