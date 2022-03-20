package auth

import "github.com/labstack/echo"

type Handler interface {
	SignUp() echo.HandlerFunc
	SignIn() echo.HandlerFunc
	// Logout() echo.HandlerFunc
}
