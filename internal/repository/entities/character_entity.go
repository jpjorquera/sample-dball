package entities

import "gorm.io/gorm"

type CharacterEntity struct {
	gorm.Model
	ExternalID  uint `gorm:"uniqueIndex"`
	Name        string
	Race        string
	Ki          string
	Description string
}

func (CharacterEntity) TableName() string {
	return "characters"
}
