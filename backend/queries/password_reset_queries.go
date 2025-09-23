package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
)

func CreatePasswordReset(req *dto.CreatePasswordResetRequest) (*models.PasswordReset, error) {
	query := `
        INSERT INTO password_resets (user_id, token, expires_at, used)
        VALUES (:user_id, :token, :expires_at, FALSE)
        RETURNING id, user_id, token, expires_at, used
    `
	var pr models.PasswordReset
	rows, err := database.DB.NamedQuery(query, req)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&pr); err != nil {
			return nil, err
		}
		return &pr, nil
	}
	return nil, err
}

func GetPasswordResetByToken(token string) (*models.PasswordReset, error) {
	var pr models.PasswordReset
	err := database.DB.Get(&pr, `
        SELECT id, user_id, token, expires_at, used
        FROM password_resets WHERE token = $1
    `, token)
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func MarkPasswordResetUsed(id string) error {
	_, err := database.DB.Exec(`UPDATE password_resets SET used = TRUE WHERE id = $1`, id)
	return err
}
