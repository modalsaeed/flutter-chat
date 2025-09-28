package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
)

func CreateEmailVerification(req *dto.CreateEmailVerificationRequest) (*models.EmailVerification, error) {
	query := `
        INSERT INTO email_verifications (user_id, token, expires_at, used)
        VALUES (:user_id, :token, :expires_at, FALSE)
        RETURNING id, user_id, token, expires_at, used
    `
	var ev models.EmailVerification
	rows, err := database.DB.NamedQuery(query, req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&ev); err != nil {
			return nil, err
		}
		return &ev, nil
	}
	return nil, err
}

func GetEmailVerification(userID, token string) (*models.EmailVerification, error) {
	var ev models.EmailVerification
	err := database.DB.Get(&ev, `
        SELECT id, user_id, token, expires_at, used
        FROM email_verifications WHERE token = $1 AND id = $2
    `, token, userID)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

func MarkEmailVerificationUsed(id string) error {
	_, err := database.DB.Exec(`UPDATE email_verifications SET used = TRUE WHERE id = $1`, id)
	return err
}
