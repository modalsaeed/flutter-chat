package models

import "time"

type MessageReceipt struct {
	ID        string    `db:"id" json:"id"`
	MessageID string    `db:"message_id" json:"message_id"`
	UserID    *string   `db:"user_id" json:"user_id,omitempty"`
	Status    string    `db:"status" json:"status"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
