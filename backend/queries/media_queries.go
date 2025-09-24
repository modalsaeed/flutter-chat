package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func CreateMedia(req *dto.CreateMediaRequest) (*models.Media, error) {
	query := `
        INSERT INTO media (uploader_id, chat_id, message_id, url, type, mime_type, size_bytes, created_at)
        VALUES (:uploader_id, :chat_id, :message_id, :url, :type, :mime_type, :size_bytes, :created_at)
        RETURNING id, uploader_id, chat_id, message_id, url, type, mime_type, size_bytes, created_at
    `
	params := map[string]interface{}{
		"uploader_id": req.UploaderID,
		"chat_id":     req.ChatID,
		"message_id":  req.MessageID,
		"url":         req.URL,
		"type":        req.Type,
		"mime_type":   req.MimeType,
		"size_bytes":  req.SizeBytes,
		"created_at":  time.Now().UTC(),
	}
	var media models.Media
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&media); err != nil {
			return nil, err
		}
		return &media, nil
	}
	return nil, err
}

func GetMediaByID(id string) (*models.Media, error) {
	var media models.Media
	err := database.DB.Get(&media, `
        SELECT id, uploader_id, chat_id, message_id, url, type, mime_type, size_bytes, created_at
        FROM media WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func ListMediaByChatID(chatID string, limit, offset int) ([]models.Media, error) {
	var media []models.Media
	err := database.DB.Select(&media, `
        SELECT id, uploader_id, chat_id, message_id, url, type, mime_type, size_bytes, created_at
        FROM media WHERE chat_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `, chatID, limit, offset)
	return media, err
}

func ListMediaByMessageID(messageID string) ([]models.Media, error) {
	var media []models.Media
	err := database.DB.Select(&media, `
        SELECT id, uploader_id, chat_id, message_id, url, type, mime_type, size_bytes, created_at
        FROM media WHERE message_id = $1
    `, messageID)
	return media, err
}

func DeleteMedia(id string) error {
	_, err := database.DB.Exec(`DELETE FROM media WHERE id = $1`, id)
	return err
}
