package dto

type RegisterRequest struct {
	Username        string  `json:"username" validate:"required,min=3,max=25"`
	Email           string  `json:"email" validate:"required,email"`
	Password        string  `json:"password" validate:"required,passwd"`
	ConfirmPassword string  `json:"confirm_password" validate:"required,eqfield=Password"`
	AvatarURL       *string `json:"avatar_url,omitempty"`
	DisplayName     *string `json:"display_name,omitempty"`
	About           *string `json:"about,omitempty"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type EditUserRequest struct {
	ID          string  `json:"id" validate:"required"`
	Username    string  `json:"username" validate:"omitempty,min=3,max=25"`
	Email       string  `json:"email" validate:"omitempty,email"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	About       *string `json:"about,omitempty"`
}

type UserResponse struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	About       *string `json:"about,omitempty"`
	CreatedAt   string  `json:"created_at"`
	LastSeen    *string `json:"last_seen,omitempty"`
}
