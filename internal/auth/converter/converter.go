package converter

import (
	"dangquang9a/go-location/internal/auth/presenter"
	"dangquang9a/go-location/internal/models"
)

func Convert_model_user_sign_up_response(user models.User) presenter.SignUpResponse {
	return presenter.SignUpResponse{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
