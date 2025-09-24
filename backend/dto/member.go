package dto

import "time"

type AddMemberRequest struct {
	ChatID string `json:"chat_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role" validate:"required"`
}

type UpdateMemberRoleRequest struct {
	ChatID string `json:"chat_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role" validate:"required"`
}

type MuteMemberRequest struct {
	ChatID     string     `json:"chat_id" validate:"required"`
	UserID     string     `json:"user_id" validate:"required"`
	MutedUntil *time.Time `json:"muted_until"`
}

type BanMemberRequest struct {
	ChatID string `json:"chat_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}
