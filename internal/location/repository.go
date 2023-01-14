package location

import (
	"context"
	"dangquang9a/go-location/internal/models"
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, todo *models.Location) error
	GetLocationsByUserID(ctx context.Context, userId string) ([]*models.Location, error)
	GetAllLocation(ctx context.Context) ([]*models.Location, error)
	CountTodo(ctx context.Context, userId string) (int, error)
}
