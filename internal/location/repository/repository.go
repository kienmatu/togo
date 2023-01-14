package repository

import (
	"context"
	location "dangquang9a/go-location/internal/location"
	"dangquang9a/go-location/internal/models"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewLocRepository(db *gorm.DB) location.LocationRepository {
	return &todoRepository{db: db}
}

func (tr *todoRepository) CreateLocation(ctx context.Context, todo *models.Location) error {
	result := tr.db.WithContext(ctx).Create(&todo)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *todoRepository) GetLocationsByUserID(ctx context.Context, userId string) ([]*models.Location, error) {
	var location []*models.Location
	err := tr.db.WithContext(ctx).Where(&models.Location{CreatedBy: userId}).Find(&location).Error

	if err != nil {
		return nil, err
	}
	return location, nil
}

func (tr *todoRepository) GetAllLocation(ctx context.Context) ([]*models.Location, error) {
	var locs []*models.Location
	// can add offset later
	err := tr.db.WithContext(ctx).Limit(200).Find(&locs).Error

	if err != nil {
		return nil, err
	}
	return locs, nil
}

func (tr *todoRepository) CountTodo(ctx context.Context, userId string) (int, error) {
	var count int
	err := tr.db.WithContext(ctx).Raw(`SELECT 
			COUNT(*)
			FROM "location"
			WHERE location.created_by = ?
			AND DATE_TRUNC('day', "created_at") = CURRENT_DATE
			GROUP BY DATE_TRUNC('day', "created_at")`, userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
