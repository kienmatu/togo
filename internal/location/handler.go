package location

import "github.com/labstack/echo/v4"

type Handler interface {
	GetAll() echo.HandlerFunc
	GetUserLocation() echo.HandlerFunc
	AddLocation() echo.HandlerFunc
}
