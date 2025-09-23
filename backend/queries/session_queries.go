package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
)

func CreateSession(req *dto.CreateSessionRequest) (*models.Session, error) {
	query := `
        INSERT INTO sessions (user_id, ip_address, user_agent, created_at, expires_at, revoked)
        VALUES (:user_id, :ip_address, :user_agent, NOW(), :expires_at, FALSE)
        RETURNING id, user_id, ip_address, user_agent, created_at, expires_at, revoked
    `
	var session models.Session
	rows, err := database.DB.NamedQuery(query, req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&session); err != nil {
			return nil, err
		}
		return &session, nil
	}
	return nil, err
}

func GetSessionByID(id string) (*models.Session, error) {
	var session models.Session
	err := database.DB.Get(&session, `
        SELECT id, user_id, ip_address, user_agent, created_at, expires_at, revoked
        FROM sessions WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func GetSessionsByUserID(userID string) ([]models.Session, error) {
	var sessions []models.Session
	err := database.DB.Select(&sessions, `
        SELECT id, user_id, ip_address, user_agent, created_at, expires_at, revoked
        FROM sessions WHERE user_id = $1
    `, userID)
	return sessions, err
}

func RevokeSession(id string) error {
	_, err := database.DB.Exec(`UPDATE sessions SET revoked = TRUE WHERE id = $1`, id)
	return err
}

func DeleteSession(id string) error {
	_, err := database.DB.Exec(`DELETE FROM sessions WHERE id = $1`, id)
	return err
}
