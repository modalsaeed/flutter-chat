package models

import "time"

type Chat struct {
	ID          string    `db:"id" json:"id"`
	Name        *string   `db:"name" json:"name,omitempty"`
	IsGroup     bool      `db:"is_group" json:"is_group"`
	CreatedBy   *string   `db:"created_by" json:"created_by,omitempty"`
	Description *string   `db:"description" json:"description,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	AvatarURL   *string   `db:"avatar_url" json:"avatar_url,omitempty"`
	IsActive    bool      `db:"is_active" json:"is_active"`
}
