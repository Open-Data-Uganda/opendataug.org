package controllers

import (
	"errors"
	"fmt"
	"strings"

	"net/mail"

	"gorm.io/gorm"
	"opendataug.org/commons"
	"opendataug.org/database"
	"opendataug.org/models"
	"opendataug.org/services"
	"opendataug.org/utils"
)

type UserController struct {
	db         *database.Database
	jwtService *services.JWTService
}

func NewUserController(db *database.Database, jwtService *services.JWTService) *UserController {
	return &UserController{
		db:         db,
		jwtService: jwtService,
	}
}

type LoginResponse struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
	UserNumber   string  `json:"user_number"`
	Role         string  `json:"role"`
	ExpiresIn    *int64  `json:"expires_in"`
}

type ResetPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type UserProfileResponse struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (c *UserController) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := c.db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func (c *UserController) FindByNumber(number string) (*models.User, error) {
	var user models.User
	if err := c.db.DB.Where("number = ?", number).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func (c *UserController) FindByAuthID(provider, authID string) (*models.User, error) {
	var user models.User
	if err := c.db.DB.Where("provider = ? AND auth_number = ?", provider, authID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return &user, nil
}

func (c *UserController) UpdateStatus(userNumber, status string) error {
	result := c.db.DB.Model(&models.User{}).Where("number = ?", userNumber).Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("failed to update status: %w", result.Error)
	}
	return nil
}

func (c *UserController) Create(user *models.User) error {
	return c.db.DB.Create(user).Error
}

func (c *UserController) GetPasswordByUserNumber(userNumber string) (*models.UserPassword, error) {
	var password models.UserPassword
	if err := c.db.DB.Where("user_number = ?", userNumber).First(&password).Error; err != nil {
		return nil, fmt.Errorf("password not found: %w", err)
	}
	return &password, nil
}

func (c *UserController) SavePasswordReset(reset *models.PasswordReset) error {
	return c.db.DB.Save(reset).Error
}

func (c *UserController) FindPasswordResetByToken(token string) (*models.PasswordReset, error) {
	var reset models.PasswordReset
	if err := c.db.DB.Where("token = ?", token).First(&reset).Error; err != nil {
		return nil, fmt.Errorf("reset token not found")
	}
	return &reset, nil
}

func (c *UserController) ExecutePasswordReset(tx *gorm.DB, userNumber string, hashedPassword string) error {
	var userPassword models.UserPassword
	if err := tx.Where("user_number = ?", userNumber).First(&userPassword).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to check existing password: %w", err)
		}
		userPassword = models.UserPassword{
			Number:       utils.UUIDGenerator(),
			UserPassword: hashedPassword,
			UserNumber:   userNumber,
		}
		return tx.Create(&userPassword).Error
	}

	return tx.Model(&userPassword).Update("user_password", hashedPassword).Error
}

func (c *UserController) AuthenticateUser(email, password string) (*models.User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, fmt.Errorf("invalid email address")
	}

	user, err := c.FindByEmail(strings.ToLower(email))
	if err != nil {
		return nil, err
	}

	if user.Status != "ACTIVE" {
		return nil, fmt.Errorf("account is not active")
	}

	userPassword, err := c.GetPasswordByUserNumber(user.Number)
	if err != nil {
		return nil, fmt.Errorf("password not set")
	}

	if _, err := commons.ComparePassword(userPassword.UserPassword, password); err != nil {
		return nil, fmt.Errorf("Invalid password or email address")
	}

	return user, nil
}

func (c *UserController) RefreshUserSession(refreshToken string) (*services.TokenDetails, error) {
	return c.jwtService.RefreshToken(refreshToken)
}

func (c *UserController) InitiatePasswordReset(email string) (*models.User, error) {
	user, err := c.FindByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("User with email address does not exist.")
	}

	return user, nil
}

func (c *UserController) SetNewPassword(token, userNumber, newPassword, confirmPassword string) error {
	if newPassword != confirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	reset, err := c.FindPasswordResetByToken(token)
	if err != nil {
		return fmt.Errorf("Password reset link has expired")
	}

	if userNumber != reset.UserNumber {
		return fmt.Errorf("invalid user number")
	}

	if reset.Status != "ACTIVE" {
		return fmt.Errorf("Password reset link has expired")
	}

	hashedPassword, err := commons.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	tx := c.db.DB.Begin()

	if err := tx.Model(&models.User{}).Where("number = ?", reset.UserNumber).Update("status", "ACTIVE").Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to activate user account")
	}

	if err := c.ExecutePasswordReset(tx, reset.UserNumber, hashedPassword); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to reset password: %w", err)
	}

	if err := tx.Model(reset).Update("status", "INACTIVE").Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update reset token: %w", err)
	}

	return tx.Commit().Error
}

func (c *UserController) InvalidateSession(refreshToken string) error {
	if refreshToken == "" {
		return fmt.Errorf("missing refresh token")
	}

	_, err := c.jwtService.ValidateToken(refreshToken)
	if err != nil {
		return fmt.Errorf("invalid refresh token: %w", err)
	}

	return nil
}

func (c *UserController) CreateLoginSession(user *models.User) (*LoginResponse, *services.TokenDetails, error) {
	tokenDetails, err := c.jwtService.CreateToken(user.Number, user.Role)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create token: %w", err)
	}

	response := &LoginResponse{
		AccessToken:  tokenDetails.AccessToken,
		RefreshToken: tokenDetails.RefreshToken,
		UserNumber:   user.Number,
		Role:         user.Role,
		ExpiresIn:    tokenDetails.AccessTokenExpiresIn,
	}

	return response, tokenDetails, nil
}

func (c *UserController) GetUserProfile(userNumber string) (*UserProfileResponse, error) {
	user, err := c.FindByNumber(userNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &UserProfileResponse{
		Email:     user.Email,
		Name:      user.FirstName + " " + user.LastName,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (c *UserController) CheckEmailExists(email string) (bool, error) {
	var user models.User
	result := c.db.DB.Where("email = ?", email).First(&user)
	return result.RowsAffected > 0, result.Error
}

func (c *UserController) CheckPhoneExists(phone string) (bool, error) {
	var user models.User
	result := c.db.DB.Where("phone = ?", phone).First(&user)
	return result.RowsAffected > 0, result.Error
}
