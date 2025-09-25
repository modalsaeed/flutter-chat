package services

import (
	"flutter-chat/dto"
	"flutter-chat/queries"
	"fmt"
	"time"
)

type AuthService struct {
	EmailService *EmailService
}

func NewAuthService(emailService *EmailService) *AuthService {
	return &AuthService{
		EmailService: emailService,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) error {
	user, err := queries.CreateUser(req)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	token, err := s.EmailService.SendVerificationEmail(user.Email)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	EmailReq := &dto.CreateEmailVerificationRequest{
		UserID:    user.ID,
		Token:     *token,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	_, err = queries.CreateEmailVerification(EmailReq)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
