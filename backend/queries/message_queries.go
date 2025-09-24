package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func CreateMessage(req *dto.CreateMessageRequest) (*models.Message, error) {
	query := `
        INSERT INTO messages (chat_id, sender_id, content, type, reply_to_id, created_at, edited, deleted)
        VALUES (:chat_id, :sender_id, :content, :type, :reply_to_id, :created_at, FALSE, FALSE)
        RETURNING id, chat_id, sender_id, content, type, reply_to_id, created_at, edited, deleted
    `
	params := map[string]interface{}{
		"chat_id":     req.ChatID,
		"sender_id":   req.SenderID,
		"content":     req.Content,
		"type":        req.Type,
		"reply_to_id": req.ReplyToID,
		"created_at":  time.Now().UTC(),
	}
	var msg models.Message
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&msg); err != nil {
			return nil, err
		}
		return &msg, nil
	}
	return nil, err
}

func GetMessageByID(id string) (*models.Message, error) {
	var msg models.Message
	err := database.DB.Get(&msg, `
        SELECT id, chat_id, sender_id, content, type, reply_to_id, created_at, edited, deleted
        FROM messages WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func ListMessagesByChatID(chatID string, limit, offset int) ([]models.Message, error) {
	var msgs []models.Message
	err := database.DB.Select(&msgs, `
        SELECT id, chat_id, sender_id, content, type, reply_to_id, created_at, edited, deleted
        FROM messages
        WHERE chat_id = $1
        ORDER BY created_at ASC
        LIMIT $2 OFFSET $3
    `, chatID, limit, offset)
	return msgs, err
}

func EditMessage(req *dto.EditMessageRequest) (*models.Message, error) {
	query := `
        UPDATE messages
        SET content = :content, edited = TRUE
        WHERE id = :id
        RETURNING id, chat_id, sender_id, content, type, reply_to_id, created_at, edited, deleted
    `
	params := map[string]interface{}{
		"id":      req.ID,
		"content": req.Content,
	}
	var msg models.Message
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&msg); err != nil {
			return nil, err
		}
		return &msg, nil
	}
	return nil, err
}

func DeleteMessage(id string) error {
	_, err := database.DB.Exec(`
        UPDATE messages SET deleted = TRUE WHERE id = $1
    `, id)
	return err
}
