package services

import (
	"flutter-chat/dto"
	"flutter-chat/queries"
	"flutter-chat/utils"
	"net/http"
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

func (s *AuthService) Register(req *dto.RegisterRequest) *dto.ErrorData {
	user, err := queries.CreateUser(req)
	if err != nil {
		return &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create user: " + err.Error(),
		}
	}

	token, err := s.EmailService.SendVerificationEmail(user.Email)
	if err != nil {
		return &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to send verification email: " + err.Error(),
		}
	}

	EmailReq := &dto.CreateEmailVerificationRequest{
		UserID:    user.ID,
		Token:     *token,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	_, err = queries.CreateEmailVerification(EmailReq)
	if err != nil {
		return &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create email verification entry: " + err.Error(),
		}
	}

	return nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.UserResponse, *dto.ErrorData) {
	user, err := queries.GetUserByIdentifier(req.Identifier)
	if err != nil {
		return nil, &dto.ErrorData{
			Code:    http.StatusUnauthorized,
			Message: "User not found",
			Fields: []dto.FieldError{
				{Field: "identifier", Message: "Invalid Email or Username"},
			},
		}
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return nil, &dto.ErrorData{
			Code:    http.StatusUnauthorized,
			Message: "Password hash check failed",
			Fields: []dto.FieldError{
				{Field: "password", Message: "Invalid password"},
			},
		}
	}

	if !user.EmailVerified {
		return nil, &dto.ErrorData{
			Code:    http.StatusUnauthorized,
			Message: "Email is unverified, verify your email to login",
		}
	}

	UserResponse := utils.ToUserResponse(user)

	return UserResponse, nil
}

func (s *AuthService) RequestPasswordReset(req *dto.PasswordResetRequest) (*string, *dto.ErrorData) {
	user, err := queries.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &dto.ErrorData{
			Code:    http.StatusNotFound,
			Message: "No user found linked to the email",
		}
	}

	token, err := s.EmailService.SendPasswordResetEmail(req.Email)
	if err != nil {
		return nil, &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to send email: " + err.Error(),
		}
	}

	resetReq := &dto.CreatePasswordReset{
		UserID:    user.ID,
		Token:     *token,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	_, err = queries.CreatePasswordReset(resetReq)
	if err != nil {
		return nil, &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create password reset entry in database: " + err.Error(),
		}
	}

	return &user.ID, nil
}

func (s *AuthService) VerifyPasswordResetOTP(req *dto.VerifyPasswordResetOTPRequest) (bool, *dto.ErrorData) {
	passwordReset, err := queries.GetPasswordReset(req.UserID, req.Token)
	if err != nil || passwordReset.Used || passwordReset.ExpiresAt.Before(time.Now()) {
		return false, &dto.ErrorData{
			Code:    http.StatusUnauthorized,
			Message: "Invalid or expired OTP",
		}
	}

	return true, nil
}

func (s *AuthService) ResetPasword(req *dto.ResetPasswordWithOTPRequest) (bool, *dto.ErrorData) {
	passwordReset, err := queries.GetPasswordReset(req.UserID, req.Token)
	if err != nil || passwordReset.Used || passwordReset.ExpiresAt.Before(time.Now()) {
		return false, &dto.ErrorData{
			Code:    http.StatusUnauthorized,
			Message: "Invalid or expired OTP",
		}
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return false, &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to hash password: " + err.Error(),
		}
	}

	err = queries.UpdateUserPassword(req.UserID, hashedPassword)
	if err != nil {
		return false, &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update password: " + err.Error(),
		}
	}

	err = queries.MarkPasswordResetUsed(passwordReset.ID)
	if err != nil {
		// Password already changed at this point so return true
		return true, &dto.ErrorData{
			Code:    http.StatusInternalServerError,
			Message: "Failed to mark password reset entry as used: " + err.Error(),
		}
	}

	return true, nil
}
