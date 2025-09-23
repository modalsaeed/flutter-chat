package models

import "time"

type Media struct {
	ID         string    `db:"id" json:"id"`
	UploaderID *string   `db:"uploader_id" json:"uploader_id,omitempty"`
	ChatID     string    `db:"chat_id" json:"chat_id"`
	MessageID  string    `db:"message_id" json:"message_id"`
	URL        string    `db:"url" json:"url"`
	Type       string    `db:"type" json:"type"`
	MimeType   *string   `db:"mime_type" json:"mime_type,omitempty"`
	SizeBytes  *int64    `db:"size_bytes" json:"size_bytes,omitempty"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
