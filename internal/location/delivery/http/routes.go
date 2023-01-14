package http

import (
	"dangquang9a/go-location/internal/location"
	"dangquang9a/go-location/internal/middlewares"

	"github.com/labstack/echo/v4"
)

// Map auth routes
func MapAuthRoutes(locGroup *echo.Group, h location.Handler, mw *middlewares.MiddlewareManager) {
	locGroup.POST("/", h.AddLocation(), mw.JWTValidation)
	locGroup.GET("/:userId", h.GetUserLocation(), mw.JWTValidation)
	locGroup.GET("/", h.GetAll(), mw.JWTValidation)
}
