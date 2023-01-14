package auth

import (
	"context"
	"dangquang9a/go-location/internal/auth/presenter"
	"dangquang9a/go-location/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, user presenter.SignUpInput) (*models.User, error)
	SignIn(ctx context.Context, username, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (string, error)
}
