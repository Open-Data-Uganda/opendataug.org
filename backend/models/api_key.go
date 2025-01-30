package models

import (
	"time"

	"gorm.io/gorm"
)

type APIKey struct {
	gorm.Model
	Number     string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	UserNumber string `gorm:"not null;index"`
	Name       string `gorm:"not null"`
	Key        string `gorm:"uniqueIndex;not null"`
	LastUsedAt *time.Time
	ExpiresAt  *time.Time
	UsageCount int64 `gorm:"default:0"`
	IsActive   bool  `gorm:"default:true"`
}
