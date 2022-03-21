package auth

import "github.com/labstack/echo/v4"

type Handler interface {
	SignUp() echo.HandlerFunc
	SignIn() echo.HandlerFunc
	// Logout() echo.HandlerFunc
}
