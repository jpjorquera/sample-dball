package repository

import (
	"dballz/internal/db"
	"dballz/internal/dto"
)

type CharacterDBRepository struct {
	db db.DB
}

func NewCharacterDBRepository(database db.DB) *CharacterDBRepository {
	return &CharacterDBRepository{db: database}
}

func (r *CharacterDBRepository) GetByName(name string) (dto.CharacterInformation, error) {
	return dto.CharacterInformation{}, nil
}

func (r *CharacterDBRepository) Save(character dto.CharacterInformation) error {
	return nil
}
