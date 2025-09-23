package models

import "time"

type User struct {
	ID            string     `db:"id" json:"id"`
	Username      string     `db:"username" json:"username"`
	Email         string     `db:"email" json:"email"`
	PasswordHash  string     `db:"password_hash" json:"-"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	LastSeen      *time.Time `db:"last_seen" json:"last_seen,omitempty"`
	AvatarURL     *string    `db:"avatar_url" json:"avatar_url,omitempty"`
	DisplayName   *string    `db:"display_name" json:"display_name,omitempty"`
	About         *string    `db:"about" json:"about,omitempty"`
	EmailVerified bool       `db:"email_verified" json:"email_verified"`
	IsActive      bool       `db:"is_active" json:"is_active"`
}
