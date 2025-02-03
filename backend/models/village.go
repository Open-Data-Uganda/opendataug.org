package models

import (
	"gorm.io/gorm"
)

type Village struct {
	Number       string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name         string `json:"name"`
	ParishNumber string `gorm:"type:varchar(36)" json:"parish_number"`
	Parish       Parish `gorm:"foreignKey:ParishNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"parish_details,omitempty"`
	gorm.Model
}

type VillageResponse struct {
	Name string `json:"name"`
}
