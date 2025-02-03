package models

import (
	"gorm.io/gorm"
)

type SubCounty struct {
	Number       string   `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name         string   `json:"name"`
	CountyNumber string   `gorm:"type:varchar(36)" json:"county_number"`
	County       County   `gorm:"foreignKey:CountyNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"county_details,omitempty"`
	Parishes     []Parish `gorm:"foreignKey:SubCountyNumber;references:Number;constraint: OnUpdate:CASCADE, OnDelete:RESTRICT;" json:"parishes,omitempty"`
	gorm.Model
}

type SubCountyResponse struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}
