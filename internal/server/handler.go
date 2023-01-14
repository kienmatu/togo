package server

import (
	authRepository "dangquang9a/go-location/internal/auth/repository"
	authUsecase "dangquang9a/go-location/internal/auth/usecase"
	locationRepository "dangquang9a/go-location/internal/location/repository"
	locationUsecase "dangquang9a/go-location/internal/location/usecase"
	"dangquang9a/go-location/internal/middlewares"

	authHttp "dangquang9a/go-location/internal/auth/delivery/http"
	locHttp "dangquang9a/go-location/internal/location/delivery/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// repos
	userRepo := authRepository.NewUserRepository(s.db)
	locRepo := locationRepository.NewLocRepository(s.db)

	//usecase
	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	locUC := locationUsecase.NewLocationUseCase(locRepo, userRepo)

	//handler
	authHandler := authHttp.NewAuthHandler(authUC)
	locHandler := locHttp.NewLocHandler(locUC)

	//middlewares
	mw := middlewares.NewMiddlewareManager(authUC)

	e.Use(middleware.BodyLimit("2M"))
	// e.Use(mw.JWTValidation())

	//versioning
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	locGroup := v1.Group("/location")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	locHttp.MapAuthRoutes(locGroup, locHandler, mw)

	return nil
}
