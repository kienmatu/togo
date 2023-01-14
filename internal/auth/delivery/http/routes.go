package http

import (
	"dangquang9a/go-location/internal/auth"

	"github.com/labstack/echo/v4"
)

// Map auth routes
func MapAuthRoutes(authGroup *echo.Group, h auth.Handler) {
	authGroup.POST("/register", h.SignUp())
	authGroup.POST("/login", h.SignIn())
	// authGroup.POST("/logout", h.Logout())

}
