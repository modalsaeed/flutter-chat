package models

import "time"

type Message struct {
	ID        string    `db:"id" json:"id"`
	ChatID    string    `db:"chat_id" json:"chat_id"`
	SenderID  *string   `db:"sender_id" json:"sender_id,omitempty"`
	Content   *string   `db:"content" json:"content,omitempty"`
	Type      string    `db:"type" json:"type"`
	ReplyToID *string   `db:"reply_to_id" json:"reply_to_id,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Edited    bool      `db:"edited" json:"edited"`
	Deleted   bool      `db:"deleted" json:"deleted"`
}
