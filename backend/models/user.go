package models

import (
	"gorm.io/gorm"
)

type User struct {
	Number    string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Email     string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"not null"`
	Provider  string `gorm:"not null"`
	AuthID    string `gorm:"uniqueIndex;not null"`
	AvatarURL string
	gorm.Model
}
