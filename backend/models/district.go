package models

import (
	"gorm.io/gorm"
)

type District struct {
	Number       string   `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name         string   `json:"name"`
	Size         float32  `json:"size"`
	TownStatus   bool     `json:"townstatus"`
	Counties     []County `gorm:"foreignKey:DistrictNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"counties,omitempty"`
	RegionNumber string   `gorm:"type:varchar(36)" json:"region_number"`
	Region       Region   `json:"region,omitempty" gorm:"foreignKey:RegionNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;"`
	gorm.Model
}

type DistrictResponse struct {
	Number       string        `json:"number"`
	Name         string        `json:"name"`
	Size         float32       `json:"size"`
	TownStatus   bool          `json:"town_status"`
	RegionNumber string        `json:"region_number"`
	Region       RegionSummary `json:"region"`
}

type RegionSummary struct {
	Name string `json:"name"`
}

type DistrictSummary struct {
	Number     string  `json:"number"`
	Name       string  `json:"name"`
	Size       float32 `json:"size"`
	TownStatus bool    `json:"town_status"`
}

type RegionWithDistricts struct {
	Number    string            `json:"number"`
	Name      string            `json:"name"`
	Districts []DistrictSummary `json:"districts"`
}
