package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func AddMember(req *dto.AddMemberRequest) error {
	_, err := database.DB.Exec(`
        INSERT INTO chat_members (chat_id, user_id, joined_at, role, is_muted, is_banned)
        VALUES ($1, $2, $3, $4, FALSE, FALSE)
        ON CONFLICT DO NOTHING
    `, req.ChatID, req.UserID, time.Now().UTC(), req.Role)
	return err
}

func UpdateMemberRole(req *dto.UpdateMemberRoleRequest) error {
	_, err := database.DB.Exec(`
        UPDATE chat_members SET role = $1 WHERE chat_id = $2 AND user_id = $3
    `, req.Role, req.ChatID, req.UserID)
	return err
}

func MuteMember(req *dto.MuteMemberRequest) error {
	_, err := database.DB.Exec(`
        UPDATE chat_members SET is_muted = TRUE, muted_until = $1 WHERE chat_id = $2 AND user_id = $3
    `, req.MutedUntil, req.ChatID, req.UserID)
	return err
}

func UnmuteMember(chatID, userID string) error {
	_, err := database.DB.Exec(`
        UPDATE chat_members SET is_muted = FALSE, muted_until = NULL WHERE chat_id = $1 AND user_id = $2
    `, chatID, userID)
	return err
}

func BanMember(req *dto.BanMemberRequest) error {
	_, err := database.DB.Exec(`
        UPDATE chat_members SET is_banned = TRUE WHERE chat_id = $1 AND user_id = $2
    `, req.ChatID, req.UserID)
	return err
}

func UnbanMember(chatID, userID string) error {
	_, err := database.DB.Exec(`
        UPDATE chat_members SET is_banned = FALSE WHERE chat_id = $1 AND user_id = $2
    `, chatID, userID)
	return err
}

func RemoveMember(chatID, userID string) error {
	_, err := database.DB.Exec(`
        DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2
    `, chatID, userID)
	return err
}

func GetMember(chatID, userID string) (*models.Member, error) {
	var member models.Member
	err := database.DB.Get(&member, `
        SELECT id, chat_id, user_id, joined_at, role, is_muted, is_banned, muted_until
        FROM chat_members WHERE chat_id = $1 AND user_id = $2
    `, chatID, userID)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func ListMembers(chatID string) ([]models.Member, error) {
	var members []models.Member
	err := database.DB.Select(&members, `
        SELECT id, chat_id, user_id, joined_at, role, is_muted, is_banned, muted_until
        FROM chat_members WHERE chat_id = $1
    `, chatID)
	return members, err
}
