package utils

import (
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func ToUserResponse(user *models.User) *dto.UserResponse {
	var lastSeen *string
	if user.LastSeen != nil {
		ls := user.LastSeen.Format(time.RFC3339)
		lastSeen = &ls
	}
	return &dto.UserResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AvatarURL:   user.AvatarURL,
		DisplayName: user.DisplayName,
		About:       user.About,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		LastSeen:    lastSeen,
	}
}
