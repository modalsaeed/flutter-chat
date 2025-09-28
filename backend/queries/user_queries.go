package queries

import (
	"flutter-chat/database"
	"flutter-chat/dto"
	"flutter-chat/models"
	"flutter-chat/utils"
	"time"
)

func CreateUser(request *dto.RegisterRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	query := `
        INSERT INTO users (username, email, password_hash, avatar_url, display_name, about)
        VALUES (:username, :email, :password_hash, :avatar_url, :display_name, :about)
        RETURNING id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
    `

	params := map[string]interface{}{
		"username":      request.Username,
		"email":         request.Email,
		"password_hash": hashedPassword,
		"avatar_url":    request.AvatarURL,
		"display_name":  request.DisplayName,
		"about":         request.About,
	}

	var user models.User
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, err
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users WHERE id = $1
    `, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users WHERE email = $1
    `, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users WHERE username = $1
    `, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func EditUser(req *dto.EditUserRequest) (*models.User, error) {

	query := `
        UPDATE users
        SET
            username = COALESCE(NULLIF(:username, ''), username),
            email = COALESCE(NULLIF(:email, ''), email),
            avatar_url = :avatar_url,
            display_name = :display_name,
            about = :about
        WHERE id = :id
        RETURNING id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
    `

	params := map[string]interface{}{
		"id":           req.ID,
		"username":     req.Username,
		"email":        req.Email,
		"avatar_url":   req.AvatarURL,
		"display_name": req.DisplayName,
		"about":        req.About,
	}

	var user models.User
	rows, err := database.DB.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, err
}

func DeactivateUser(id string) error {
	_, err := database.DB.Exec(`UPDATE users SET is_active = FALSE WHERE id = $1`, id)
	return err
}

func ListUsers(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := database.DB.Select(&users, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CountUsers() (int, error) {
	var count int
	err := database.DB.Get(&count, `SELECT COUNT(*) FROM users`)
	return count, err
}

func SearchUsers(queryStr string, limit int) ([]models.User, error) {
	var users []models.User
	search := "%" + queryStr + "%"
	err := database.DB.Select(&users, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users
        WHERE username ILIKE $1 OR email ILIKE $1
        LIMIT $2
    `, search, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByIdentifier(identifier string) (*models.User, error) {
	var user models.User
	err := database.DB.Get(&user, `
        SELECT id, username, email, created_at, updated_at, last_seen, avatar_url, display_name, about, email_verified, is_active
        FROM users
        WHERE email = $1 OR username = $1
        LIMIT 1
    `, identifier)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateLastSeen(userID string) error {
	_, err := database.DB.Exec(`
        UPDATE users SET last_seen = $1 WHERE id = $2
    `, time.Now().UTC(), userID)
	return err
}

func ReactivateUser(id string) error {
	_, err := database.DB.Exec(`UPDATE users SET is_active = TRUE WHERE id = $1`, id)
	return err
}

func UpdateUserPassword(userID, hashedPassword string) error {
	_, err := database.DB.Exec(`
		UPDATE users SET password_hash = $1 WHERE id = $2
	`, hashedPassword, userID)
	return err
}
