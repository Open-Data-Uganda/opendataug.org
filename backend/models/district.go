package models

import (
	"gorm.io/gorm"
)

type District struct {
	Number       string   `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name         string   `json:"name"`
	TownStatus   bool     `json:"town_status"`
	Counties     []County `gorm:"foreignKey:DistrictNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"counties,omitempty"`
	RegionNumber string   `gorm:"type:varchar(36)" json:"region_number"`
	Region       Region   `json:"region,omitempty" gorm:"foreignKey:RegionNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;"`
	gorm.Model
}

type DistrictResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	TownStatus bool   `json:"town_status"`
	RegionID   string `json:"region_id"`
	RegionName string `json:"region_name"`
}
