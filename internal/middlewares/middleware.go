package middlewares

import (
	"kienmatu/go-todos/config"
	"kienmatu/go-todos/internal/auth"

	"github.com/sirupsen/logrus"
)

type MiddlewareManager struct {
	authUC auth.UseCase
	cfg    *config.Configuration
	logger *logrus.Logger
	// origins []string
}

func NewMiddlewareManager(authUC auth.UseCase, cfg *config.Configuration, logger *logrus.Logger) *MiddlewareManager {
	return &MiddlewareManager{authUC, cfg, logger}
}
