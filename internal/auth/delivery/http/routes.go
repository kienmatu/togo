package http

import (
	"kienmatu/go-todos/internal/auth"

	"github.com/labstack/echo"
)

// Map auth routes
func MapAuthRoutes(authGroup *echo.Group, h auth.Handler) {
	authGroup.POST("/register", h.SignUp())
	authGroup.POST("/login", h.SignIn())
	// authGroup.POST("/logout", h.Logout())

}
