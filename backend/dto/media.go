package dto

type CreateMediaRequest struct {
	UploaderID *string `json:"uploader_id,omitempty"`
	ChatID     string  `json:"chat_id" validate:"required"`
	MessageID  string  `json:"message_id" validate:"required"`
	URL        string  `json:"url" validate:"required"`
	Type       string  `json:"type" validate:"required"`
	MimeType   *string `json:"mime_type,omitempty"`
	SizeBytes  *int64  `json:"size_bytes,omitempty"`
}

type EditMediaRequest struct {
	ID        string  `json:"id" validate:"required"`
	URL       *string `json:"url,omitempty"`
	Type      *string `json:"type,omitempty"`
	MimeType  *string `json:"mime_type,omitempty"`
	SizeBytes *int64  `json:"size_bytes,omitempty"`
}
