package server

import (
	"kienmatu/go-todos/internal/auth/repository"
	"kienmatu/go-todos/internal/auth/usecase"

	authHttp "kienmatu/go-todos/internal/auth/delivery/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// repos
	userRepo := repository.NewUserRepository(s.db)

	//usecase
	authUC := usecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)

	//handler
	authHandler := authHttp.NewAuthHandler(authUC)

	e.Use(middleware.BodyLimit("2M"))

	//versioning
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandler)

	return nil
}
