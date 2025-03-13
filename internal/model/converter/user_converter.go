package converter

import (
	"fp-academya-be/internal/entity"
	"fp-academya-be/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		IsPremium: user.IsPremium,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}
