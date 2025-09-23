package dto

import "time"

type CreatePasswordResetRequest struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

type ResetPasswordRequest struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}
