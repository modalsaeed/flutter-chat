package dto

import (
	"time"
)

type CreateSessionRequest struct {
	UserID    string
	IPAddress string
	UserAgent string
	ExpiresAt time.Time
}
