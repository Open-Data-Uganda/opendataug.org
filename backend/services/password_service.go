package services

import (
	"fmt"
	"net/mail"
	"strings"

	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/models"
)

type PasswordService struct {
	db *gorm.DB
}

func NewPasswordService(db *gorm.DB) *PasswordService {
	return &PasswordService{
		db: db,
	}
}

func (s *PasswordService) InitiateReset(email string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if _, err := mail.ParseAddress(email); err != nil {
		return "", fmt.Errorf("invalid email address")
	}

	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	status := "INACTIVE"
	statusActive := "ACTIVE"
	var userToken string

	if user.Status == statusActive {
		if err := s.db.Model(&user).Update("status", status).Error; err != nil {
			return "", err
		}

		userToken = commons.UUIDGenerator()
		resetToken := models.PasswordReset{
			UserNumber: user.Number,
			Number:     commons.UUIDGenerator(),
			Token:      userToken,
			Status:     status,
		}

		if err := s.db.Create(&resetToken).Error; err != nil {
			return "", err
		}
	} else {
		var resetToken models.PasswordReset
		if err := s.db.Where("token = ?", userToken).First(&resetToken).Error; err != nil {
			return "", err
		}
		userToken = resetToken.Token
	}

	return userToken, nil
}

func (s *PasswordService) ResetPassword(token, newPassword, confirmPassword string) error {
	if newPassword != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	if len(newPassword) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	}

	var resetToken models.PasswordReset
	if err := s.db.Where("token = ?", token).First(&resetToken).Error; err != nil {
		return err
	}

	if resetToken.Status == "ACTIVE" {
		return fmt.Errorf("token has expired")
	}

	hashedPassword, err := commons.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.User{}).Where("number = ?", resetToken.UserNumber).
			Update("password", hashedPassword).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.User{}).Where("number = ?", resetToken.UserNumber).
			Update("status", "ACTIVE").Error; err != nil {
			return fmt.Errorf("failed to update user status: %w", err)
		}

		if err := tx.Model(&resetToken).Update("status", "ACTIVE").Error; err != nil {
			return fmt.Errorf("failed to update token status: %w", err)
		}

		return nil
	})
}
