package models

import (
	"errors"

	"gorm.io/gorm"
)

type (
	UserPassword struct {
		gorm.Model
		Number       string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
		UserPassword string `json:"-" gorm:"not null;"`
		UserNumber   string `json:"user_number" gorm:"size:36;not null"`
		User         User   `gorm:"foreignKey:UserNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;"`
	}

	UserPasswordList struct {
		Number       string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
		UserNumber   string `json:"user_number"`
		UserPassword string `json:"user_password"`
	}
)

type ResetPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (p *ResetPassword) Validate() error {
	if p.Password == "" {
		return errors.New("password is required")
	}
	if p.ConfirmPassword == "" {
		return errors.New("confirm password is required")
	}

	return nil
}
