package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func CreateChat(req *dto.CreateChatRequest) (*models.Chat, error) {
	query := `
        INSERT INTO chats (name, is_group, created_by, description, avatar_url, created_at, updated_at, is_active)
        VALUES (:name, :is_group, :created_by, :description, :avatar_url, :created_at, :updated_at, TRUE)
        RETURNING id, name, is_group, created_by, description, created_at, updated_at, avatar_url, is_active
    `
	params := map[string]interface{}{
		"name":        req.Name,
		"is_group":    req.IsGroup,
		"created_by":  req.CreatedBy,
		"description": req.Description,
		"avatar_url":  req.AvatarURL,
		"created_at":  time.Now().UTC(),
		"updated_at":  time.Now().UTC(),
	}
	var chat models.Chat
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&chat); err != nil {
			return nil, err
		}
		return &chat, nil
	}
	return nil, err
}

func GetChatByID(id string) (*models.Chat, error) {
	var chat models.Chat
	err := database.DB.Get(&chat, `
        SELECT id, name, is_group, created_by, description, created_at, updated_at, avatar_url, is_active
        FROM chats WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func EditChat(req *dto.EditChatRequest) (*models.Chat, error) {
	query := `
        UPDATE chats
        SET
            name = COALESCE(:name, name),
            description = COALESCE(:description, description),
            avatar_url = COALESCE(:avatar_url, avatar_url),
            is_active = COALESCE(:is_active, is_active),
            updated_at = :updated_at
        WHERE id = :id
        RETURNING id, name, is_group, created_by, description, created_at, updated_at, avatar_url, is_active
    `
	params := map[string]interface{}{
		"id":          req.ID,
		"name":        req.Name,
		"description": req.Description,
		"avatar_url":  req.AvatarURL,
		"is_active":   req.IsActive,
		"updated_at":  time.Now().UTC(),
	}
	var chat models.Chat
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&chat); err != nil {
			return nil, err
		}
		return &chat, nil
	}
	return nil, err
}

func DeleteChat(id string) error {
	_, err := database.DB.Exec(`DELETE FROM chats WHERE id = $1`, id)
	return err
}

func ListChats(limit, offset int) ([]models.Chat, error) {
	var chats []models.Chat
	err := database.DB.Select(&chats, `
        SELECT id, name, is_group, created_by, description, created_at, updated_at, avatar_url, is_active
        FROM chats
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)
	return chats, err
}

func AddUserToChat(chatID, userID string) error {
	_, err := database.DB.Exec(`
        INSERT INTO chat_members (chat_id, user_id) VALUES ($1, $2)
        ON CONFLICT DO NOTHING
    `, chatID, userID)
	return err
}

func RemoveUserFromChat(chatID, userID string) error {
	_, err := database.DB.Exec(`
        DELETE FROM chat_members WHERE chat_id = $1 AND user_id = $2
    `, chatID, userID)
	return err
}
