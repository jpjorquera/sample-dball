package repository

import (
	"dballz/internal/dto"

	"gorm.io/gorm"
)

type CharacterDBRepository struct {
	db *gorm.DB
}

func NewCharacterDBRepository(database *gorm.DB) *CharacterDBRepository {
	return &CharacterDBRepository{db: database}
}

func (r *CharacterDBRepository) GetByName(name string) (dto.CharacterInformation, error) {
	return dto.CharacterInformation{}, nil
}

func (r *CharacterDBRepository) Save(character dto.CharacterInformation) error {
	return nil
}
