package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	ExternalID  uint `gorm:"uniqueIndex"`
	Name        string
	Race        string
	Ki          string
	Description string
}
