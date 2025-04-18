package models

import (
	"errors"
	"fmt"
	"html"
	"strings"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "ADMIN"
	RoleUser  UserRole = "USER"
)

type User struct {
	Number    string   `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Email     string   `gorm:"uniqueIndex;not null" json:"email"`
	Name      string   `gorm:"not null" json:"name"`
	FirstName string   `gorm:"type:text;size:255;not null;" json:"first_name"`
	LastName  string   `gorm:"type:text;size:255;" json:"last_name"`
	Role      UserRole `gorm:"type:text;size:100;default:USER;" json:"role"`
	Status    string   `json:"status" gorm:"size:100;not null"`
	gorm.Model
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) ValidateRole() error {
	switch u.Role {
	case RoleAdmin, RoleUser:
		return nil
	default:
		return errors.New("invalid role")
	}
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrEmailRequired struct {
	Field string
}

func (e *ErrEmailRequired) Error() string {
	return fmt.Sprintf("%s is required", e.Field)
}

func (input *SignInRequest) Validate() error {
	if input.Email == "" {
		return &ErrEmailRequired{Field: "Email"}
	}

	if input.Password == "" {
		return &ErrEmailRequired{Field: "Password"}
	}

	return nil
}

func (input *SignInRequest) Prepare() {
	input.Email = html.EscapeString(strings.TrimSpace(strings.ToLower(input.Email)))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":

		if u.FirstName == "" {
			return errors.New("first name is required")
		}
		if u.LastName == "" {
			return errors.New("last name is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email is invalid")
		}

		return nil
	case "login":
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email is invalid")
		}

		return nil

	default:
		if u.FirstName == "" {
			return errors.New("first name is required")
		}
		if u.Status == "" {
			return errors.New("status is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("email is invalid")
		}

		return nil
	}
}

type (
	PasswordReset struct {
		gorm.Model
		Number     string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
		Token      string `json:"token" gorm:"not null"`
		Status     string `json:"status" gorm:"size:10;not null"`
		UserNumber string `json:"user_number" gorm:"size:36;index;not null"`
		User       User   `gorm:"foreignKey:UserNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;"`
	}

	PasswordResetInfo struct {
		Number     string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
		UserNumber string `json:"user_number"`
		Token      string `json:"token"`
		Status     string `json:"status"`
	}
)

type SignUpInput struct {
	FirstName string   `json:"first_name" binding:"required"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email" binding:"required"`
	Role      UserRole `json:"role"`
}

func (s *SignUpInput) Prepare() {
	s.FirstName = strings.TrimSpace(s.FirstName)
	s.LastName = strings.TrimSpace(s.LastName)
	s.Email = strings.TrimSpace(strings.ToLower(s.Email))
}

func (s *SignUpInput) Validate() error {
	if s.FirstName == "" {
		return errors.New("first name is required")
	}
	if s.Email == "" {
		return errors.New("email is required")
	}
	if s.Role != "" {
		switch s.Role {
		case RoleAdmin, RoleUser:
			return nil
		default:
			return errors.New("invalid role")
		}
	}
	return nil
}
