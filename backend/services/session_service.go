package services

import (
	"flutter-chat/dto"
	"flutter-chat/queries"
	"fmt"
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
