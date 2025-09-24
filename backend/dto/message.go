package dto

type CreateMessageRequest struct {
	ChatID    string  `json:"chat_id" validate:"required"`
	SenderID  *string `json:"sender_id,omitempty"`
	Content   *string `json:"content,omitempty"`
	Type      string  `json:"type" validate:"required"`
	ReplyToID *string `json:"reply_to_id,omitempty"`
}

type EditMessageRequest struct {
	ID      string  `json:"id" validate:"required"`
	Content *string `json:"content,omitempty"`
}

type DeleteMessageRequest struct {
	ID string `json:"id" validate:"required"`
}
