package services

import (
	"flutter-chat/dto"
	"flutter-chat/queries"
	"fmt"
	"net/http"
)

type SessionService struct {
}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func CreateSession(req *dto.CreateSessionRequest) (*string, error) {
	session, err := queries.CreateSession(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	return &session.ID, nil
}

func (s *SessionService) RevokeSession(token string) *dto.ErrorData {
	session, err := queries.GetSessionByID(token)
	if err != nil {
		return &dto.ErrorData{
			Code:    http.StatusNotFound,
			Message: "Session not found",
		}
	}

	if session.Revoked {
		return &dto.ErrorData{
			Code:    http.StatusBadRequest,
			Message: "Session already revoked",
		}
	}

	err = queries.RevokeSession(session.ID)
	if err != nil {
		return &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to revoke session: " + err.Error(),
		}
	}
	return nil
}
