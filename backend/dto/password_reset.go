package dto

import "time"

type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type CreatePasswordReset struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

type VerifyPasswordResetOTPRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Token  string `json:"otp" validate:"required"`
}

type ResetPasswordWithOTPRequest struct {
	UserID          string `json:"user_id" validate:"required"`
	Token           string `json:"otp" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,passwd"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}
