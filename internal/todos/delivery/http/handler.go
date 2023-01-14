package http

import (
	"dangquang9a/go-location/internal/auth"
	"dangquang9a/go-location/internal/models"
	"dangquang9a/go-location/internal/todos"
	"dangquang9a/go-location/internal/todos/presenter"
	"dangquang9a/go-location/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	todoUC todos.UseCase
}

func NewTodoHandler(todoUC todos.UseCase) *todoHandler {
	return &todoHandler{todoUC: todoUC}
}

func (th *todoHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := th.todoUC.GetAllTodos(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, mapTodos(todos))
	}
}

// Need to implement the permission before getting todos of another
func (th *todoHandler) GetUserTodos() echo.HandlerFunc {
	return func(c echo.Context) error {
		rawId := c.Param(auth.CtxUserKey)
		userId, err := uuid.Parse(rawId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		todos, err := th.todoUC.GetTodosByUserId(c.Request().Context(), userId.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, mapTodos(todos))
	}
}

func (th *todoHandler) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get(auth.CtxUserKey)
		input := &presenter.TodoRequest{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		err := th.todoUC.CreateTodo(c.Request().Context(), fmt.Sprintf("%v", userId), input.Note)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

func mapTodos(td []*models.Location) []*presenter.TodoResponse {
	out := make([]*presenter.TodoResponse, len(td))

	for i, b := range td {
		out[i] = mapTodo(b)
	}

	return out
}

func mapTodo(t *models.Location) *presenter.TodoResponse {
	return &presenter.TodoResponse{
		Id:        t.Id,
		Content:   t.Name,
		CreatedAt: t.CreatedAt,
		CreatedBy: t.CreatedBy,
	}
}
