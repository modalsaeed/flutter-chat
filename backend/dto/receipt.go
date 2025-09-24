package dto

type CreateReceiptRequest struct {
	MessageID string  `json:"message_id" validate:"required"`
	UserID    *string `json:"user_id,omitempty"`
	Status    string  `json:"status" validate:"required"`
}

type UpdateReceiptRequest struct {
	ID     string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}
