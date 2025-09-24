package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"time"
)

func CreateReceipt(req *dto.CreateReceiptRequest) (*models.MessageReceipt, error) {
	query := `
        INSERT INTO message_receipts (message_id, user_id, status, updated_at)
        VALUES (:message_id, :user_id, :status, :updated_at)
        RETURNING id, message_id, user_id, status, updated_at
    `
	params := map[string]interface{}{
		"message_id": req.MessageID,
		"user_id":    req.UserID,
		"status":     req.Status,
		"updated_at": time.Now().UTC(),
	}
	var receipt models.MessageReceipt
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&receipt); err != nil {
			return nil, err
		}
		return &receipt, nil
	}
	return nil, err
}

func UpdateReceipt(req *dto.UpdateReceiptRequest) (*models.MessageReceipt, error) {
	query := `
        UPDATE message_receipts
        SET status = :status, updated_at = :updated_at
        WHERE id = :id
        RETURNING id, message_id, user_id, status, updated_at
    `
	params := map[string]interface{}{
		"id":         req.ID,
		"status":     req.Status,
		"updated_at": time.Now().UTC(),
	}
	var receipt models.MessageReceipt
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&receipt); err != nil {
			return nil, err
		}
		return &receipt, nil
	}
	return nil, err
}

func GetReceiptsByMessageID(messageID string) ([]models.MessageReceipt, error) {
	var receipts []models.MessageReceipt
	err := database.DB.Select(&receipts, `
        SELECT id, message_id, user_id, status, updated_at
        FROM message_receipts
        WHERE message_id = $1
    `, messageID)
	return receipts, err
}

func GetReceiptsByUserID(userID string) ([]models.MessageReceipt, error) {
	var receipts []models.MessageReceipt
	err := database.DB.Select(&receipts, `
        SELECT id, message_id, user_id, status, updated_at
        FROM message_receipts
        WHERE user_id = $1
    `, userID)
	return receipts, err
}
