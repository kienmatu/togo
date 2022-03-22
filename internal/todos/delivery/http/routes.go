package http

import (
	"kienmatu/go-todos/internal/middlewares"
	"kienmatu/go-todos/internal/todos"

	"github.com/labstack/echo/v4"
)

// Map auth routes
func MapAuthRoutes(todoGroup *echo.Group, h todos.Handler, mw *middlewares.MiddlewareManager) {
	todoGroup.POST("/", h.AddTodo(), mw.JWTValidation)
	todoGroup.GET("/:userId", h.GetUserTodos(), mw.JWTValidation)
	todoGroup.GET("/", h.GetAll())
}
