package dto

import (
	"time"
)

type CreateEmailVerificationRequest struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
}

type VerifyEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
	Token string `json:"token" validate:"required"`
}
