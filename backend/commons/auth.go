package commons

import (
	"fmt"
	"net/mail"
	"strings"

	"gorm.io/gorm"
	"opendataug.org/models"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

type LoginResponse struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	UserNumber   string  `json:"user_number"`
	Role         string  `json:"role"`
	ExpiresIn    *int64  `json:"expires_in"`
}

type TokenDetails struct {
	AccessToken           *string
	RefreshToken          *string
	AccessTokenExpiresIn  *int64
	RefreshTokenExpiresIn *int64
}

func (s *AuthService) AuthenticateUser(email, password string) (*models.User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("invalid email address")
	}

	var user models.User
	if err := s.db.Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Status != "ACTIVE" {
		return nil, fmt.Errorf("account is not active")
	}

	var userPassword models.UserPassword
	if err := s.db.Where("user_number = ?", user.Number).First(&userPassword).Error; err != nil {
		return nil, fmt.Errorf("password not set")
	}

	if _, err := ComparePassword(userPassword.UserPassword, password); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &user, nil
}

func (s *AuthService) CreateLoginSession(user *models.User) (*LoginResponse, *TokenDetails, error) {
	tokenDetails, err := CreateToken(user.Number)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token: %w", err)
	}

	response := &LoginResponse{
		RefreshToken: tokenDetails.RefreshToken,
		AccessToken:  tokenDetails.AccessToken,
		UserNumber:   user.Number,
		Role:         user.Role,
		ExpiresIn:    tokenDetails.AccessTokenExpiresIn,
	}

	return response, tokenDetails, nil
}

func (s *AuthService) RefreshUserSession(refreshToken string) (*TokenDetails, error) {
	return RefreshToken(refreshToken)
}

func (s *AuthService) InvalidateSession(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("missing refresh token")
	}

	_, err := ValidateToken(refreshToken, "refresh")
	if err != nil {
		return fmt.Errorf("invalid refresh token: %w", err)
	}

	return nil
}
