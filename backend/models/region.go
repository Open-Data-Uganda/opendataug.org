package models

import "gorm.io/gorm"

type Region struct {
	gorm.Model
	Number    string     `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name      string     `json:"name"`
	Districts []District `json:"districts,omitempty" gorm:"foreignKey:RegionNumber;references:Number"`
}

type RegionResponse struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}
