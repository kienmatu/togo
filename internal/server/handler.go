package server

import (
	authRepository "kienmatu/go-todos/internal/auth/repository"
	authUsecase "kienmatu/go-todos/internal/auth/usecase"
	"kienmatu/go-todos/internal/middlewares"
	todoRepository "kienmatu/go-todos/internal/todos/repository"
	todoUsecase "kienmatu/go-todos/internal/todos/usecase"

	authHttp "kienmatu/go-todos/internal/auth/delivery/http"
	todoHttp "kienmatu/go-todos/internal/todos/delivery/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// repos
	userRepo := authRepository.NewUserRepository(s.db)
	todoRepo := todoRepository.NewTodoRepository(s.db)

	//usecase
	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	todoUC := todoUsecase.NewTodoUseCase(todoRepo, userRepo)

	//handler
	authHandler := authHttp.NewAuthHandler(authUC)
	todoHandler := todoHttp.NewTodoHandler(todoUC)

	//middlewares
	mw := middlewares.NewMiddlewareManager(authUC)

	e.Use(middleware.BodyLimit("2M"))
	// e.Use(mw.JWTValidation())

	//versioning
	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	todoGroup := v1.Group("/todos")

	authHttp.MapAuthRoutes(authGroup, authHandler)
	todoHttp.MapAuthRoutes(todoGroup, todoHandler, mw)

	return nil
}
