package todos

import "github.com/labstack/echo/v4"

type Handler interface {
	GetAll() echo.HandlerFunc
	GetUserTodos() echo.HandlerFunc
	AddTodo() echo.HandlerFunc
}
