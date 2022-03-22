package middlewares

import (
	"kienmatu/go-todos/internal/auth"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (mw *MiddlewareManager) JWTValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		if headerParts[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		userId, err := mw.authUC.ParseToken(c.Request().Context(), headerParts[1])
		if err != nil {
			status := http.StatusInternalServerError
			if err == auth.ErrInvalidAccessToken {
				status = http.StatusUnauthorized
			}

			return echo.NewHTTPError(status)
		}

		c.Set(auth.CtxUserKey, userId)
		return next(c)
	}
}
