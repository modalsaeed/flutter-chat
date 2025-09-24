package dto

type CreateChatRequest struct {
	Name        *string `json:"name,omitempty"`
	IsGroup     bool    `json:"is_group"`
	CreatedBy   *string `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
}

type EditChatRequest struct {
	ID          string  `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
}
