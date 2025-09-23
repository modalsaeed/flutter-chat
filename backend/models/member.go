package models

import "time"

type Member struct {
	ID         string     `db:"id" json:"id"`
	ChatID     string     `db:"chat_id" json:"chat_id"`
	UserID     string     `db:"user_id" json:"user_id"`
	JoinedAt   time.Time  `db:"joined_at" json:"joined_at"`
	Role       string     `db:"role" json:"role"`
	IsMuted    bool       `db:"is_muted" json:"is_muted"`
	IsBanned   bool       `db:"is_banned" json:"is_banned"`
	MutedUntil *time.Time `db:"muted_until" json:"muted_until,omitempty"`
}
