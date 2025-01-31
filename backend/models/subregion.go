package models

import (
	"errors"
	"html"
	"strings"

	"gorm.io/gorm"
)

type SubRegion struct {
	gorm.Model
	Number       string `gorm:"primaryKey;type:varchar(36);not null;unique" json:"number"`
	Name         string `json:"name"`
	RegionNumber string `gorm:"type:varchar(36);not null;index" json:"region_number"`
	Region       Region `gorm:"foreignKey:RegionNumber;references:Number;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"region,omitempty"`
}

func (s *SubRegion) Prepare() {
	s.Name = html.EscapeString(strings.TrimSpace(strings.ToLower(s.Name)))
}

func (s *SubRegion) Validate(action string) error {
	if s.Name == "" {
		return errors.New("sub region name is required")
	}

	return nil
}
