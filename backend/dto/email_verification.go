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
	Token string `json:"token"`
}
