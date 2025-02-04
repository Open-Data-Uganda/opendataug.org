package models

import (
	"gorm.io/gorm"
)

type County struct {
	Number         string      `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name           string      `json:"name"`
	DistrictNumber string      `gorm:"type:varchar(36)" json:"district_number"`
	District       District    `gorm:"foreignKey:DistrictNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"district_details,omitempty"`
	SubCounties    []SubCounty `gorm:"foreignKey:CountyNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"sub_counties,omitempty"`
	gorm.Model
}

type CountyResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	DistrictNumber string `json:"district_number"`
	DistrictName   string `json:"district_name"`
}
