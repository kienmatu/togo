package repository

import (
	"context"
	"kienmatu/go-todos/internal/auth"
	"kienmatu/go-todos/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db     *gorm.DB
}

func NewUserRepository(db *gorm.DB) auth.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := ur.db.Create(&user)
	
	if(result.Error != nil){
		return result.Error
	}
	return nil
}


func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := ur.db.Where(&models.User{
		Username: username,
		}).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}