package http

import (
	"dangquang9a/go-location/internal/auth"
	"dangquang9a/go-location/internal/location"
	"dangquang9a/go-location/internal/location/presenter"
	"dangquang9a/go-location/internal/models"
	"dangquang9a/go-location/utils"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	todoUC location.UseCase
}

func NewLocHandler(todoUC location.UseCase) *todoHandler {
	return &todoHandler{todoUC: todoUC}
}

// Location godoc
// @Summary Get all location in db
// @Description Get all location in db
// @Tags Location
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []presenter.CreateLocationResponse
// @Failure 401 {object} error
// @Router /location/ [get]
func (th *todoHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		location, err := th.todoUC.GetAllLocations(c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, mapTodos(location))
	}
}

// Location godoc
// @Summary Get location of current user
// @Description Provice current user's location history
// @Tags Location
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} []presenter.CreateLocationResponse
// @Failure 401 {object} error
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Router /location/{user_id} [get]
// TODO: Need to implement the permission before getting location of another, use redis
func (th *todoHandler) GetUserLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		rawId := c.Param(auth.CtxUserKey)
		userId, err := uuid.Parse(rawId)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		location, err := th.todoUC.GetLocationsByUserID(c.Request().Context(), userId.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, mapTodos(location))
	}
}

// Location godoc
// @Summary Add Location
// @Description Create new location in database
// @Tags Location
// @Accept json
// @Produce json
// @Param body body presenter.CreateLocationRequest true "Location detail"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 201 {object} string
// @Failure 401 {object} error
// @Failure 422 {object} error
// @Router /location/ [post]
func (th *todoHandler) AddLocation() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get(auth.CtxUserKey)
		input := &presenter.CreateLocationRequest{}
		if err := utils.ReadRequest(c, input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		if input.Latitude == 0 || input.Longitude == 0 {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Latitude and Longtitude cannot be zero")
		}

		err := th.todoUC.CreateLocation(c.Request().Context(), fmt.Sprintf("%v", userId), *input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, nil)
	}
}

func mapTodos(td []*models.Location) []*presenter.CreateLocationResponse {
	out := make([]*presenter.CreateLocationResponse, len(td))

	for i, b := range td {
		out[i] = convertLocation(b)
	}

	return out
}

func convertLocation(t *models.Location) *presenter.CreateLocationResponse {
	return &presenter.CreateLocationResponse{
		Id:        t.Id,
		Name:      t.Name,
		CreatedAt: t.CreatedAt,
		CreatedBy: t.CreatedBy,
		Latitude:  t.Lat,
		Longitude: t.Lng,
	}
}
