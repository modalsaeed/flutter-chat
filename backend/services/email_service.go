package services

import (
	"context"
	"flutter-chat/utils"
	"fmt"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

type EmailService struct {
	Domain string
	APIKey string
	From   string
}

func NewEmailService(domain, apiKey, from string) *EmailService {
	return &EmailService{
		Domain: domain,
		APIKey: apiKey,
		From:   from,
	}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	mg := mailgun.NewMailgun(s.Domain, s.APIKey)
	message := mailgun.NewMessage(
		s.From,
		subject,
		body,
		to,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := mg.Send(ctx, message)
	return err
}

func (s *EmailService) SendVerificationEmail(to string) error {
	token, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("Failed to generate OTP: %w", err.Error())
	}

	subject := "Verify your Email"
	body := fmt.Sprintf(
		"Hello,\n\nYour verification code is: %s\n\nEnter this code in the app to verify your email address. This code will expire in 10 minutes.\n\nIf you did not request this, please ignore this email.\n\nThanks,\nThe Flutter Chat Team",
		token,
	)

	err = s.SendEmail(to, subject, body)
	return err
}

func (s *EmailService) SendPasswordResetEmail(to string) error {
	token, err := utils.GenerateOTP()
	if err != nil {
		return fmt.Errorf("Failed to generate OTP: %w", err)
	}

	subject := "Reset Your Password"
	body := fmt.Sprintf(
		"Hello,\n\nYour password reset code is: %s\n\nEnter this code in the app to reset your password. This code will expire in 10 minutes.\n\nIf you did not request a password reset, please ignore this email.\n\nThanks,\nThe Flutter Chat Team",
		token,
	)

	err = s.SendEmail(to, subject, body)
	return err
}
