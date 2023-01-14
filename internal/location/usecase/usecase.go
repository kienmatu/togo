package usecase

import (
	"context"
	"dangquang9a/go-location/internal/auth"
	"dangquang9a/go-location/internal/location"
	"dangquang9a/go-location/internal/location/presenter"
	"dangquang9a/go-location/internal/models"
	"time"

	"github.com/google/uuid"
)

type todoUsecase struct {
	locationRepo location.LocationRepository
	userRepo     auth.UserRepository
}

func NewLocationUseCase(locationRepo location.LocationRepository, userRepo auth.UserRepository) location.UseCase {
	return &todoUsecase{
		locationRepo: locationRepo,
		userRepo:     userRepo,
	}
}

func (tu todoUsecase) CreateLocation(ctx context.Context, userId string, location presenter.CreateLocationRequest) error {
	newLoc := &models.Location{
		Id:        uuid.New().String(),
		Lat:       location.Latitude,
		Lng:       location.Longitude,
		Name:      location.Name,
		CreatedAt: time.Now(),
		CreatedBy: userId,
		Note:      location.Note,
	}

	return tu.locationRepo.CreateLocation(ctx, newLoc)

}

func (tu todoUsecase) GetLocationsByUserID(ctx context.Context, userId string) ([]*models.Location, error) {
	locations, err := tu.locationRepo.GetLocationsByUserID(ctx, userId)

	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (tu todoUsecase) GetAllLocations(ctx context.Context) ([]*models.Location, error) {
	locations, err := tu.locationRepo.GetAllLocation(ctx)

	if err != nil {
		return nil, err
	}
	return locations, nil
}
