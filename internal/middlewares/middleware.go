package middlewares

import (
	"dangquang9a/go-location/internal/auth"
)

type MiddlewareManager struct {
	authUC auth.UseCase
	// cfg    *config.Configuration
	// logger *logrus.Logger
	// origins []string
}

func NewMiddlewareManager(authUC auth.UseCase) *MiddlewareManager {
	return &MiddlewareManager{authUC}
}
