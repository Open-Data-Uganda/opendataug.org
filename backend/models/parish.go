package models

import (
	"gorm.io/gorm"
)

type Parish struct {
	Number          string    `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name            string    `json:"name"`
	SubCountyNumber string    `gorm:"type:varchar(36)" json:"subcounty_number"`
	SubCounty       SubCounty `gorm:"foreignKey:SubCountyNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"subcounty_details,omitempty"`
	Villages        []Village `gorm:"foreignKey:ParishNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"villages,omitempty"`
	gorm.Model
}

type ParishResponse struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}
