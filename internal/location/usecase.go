package location

import (
	"context"
	"dangquang9a/go-location/internal/location/presenter"
	"dangquang9a/go-location/internal/models"
)

type UseCase interface {
	CreateLocation(ctx context.Context, userId string, location presenter.CreateLocationRequest) error
	GetLocationsByUserID(ctx context.Context, userId string) ([]*models.Location, error)
	GetAllLocations(ctx context.Context) ([]*models.Location, error)
}
