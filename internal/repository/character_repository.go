package repository

import (
	"dballz/internal/db"
)

type CharacterRepository struct {
	db db.DB
}

func NewCharacterRepository(database db.DB) *CharacterRepository {
	return &CharacterRepository{db: database}
}

func (r *CharacterRepository) CreateCharacter() (int, error) {
	return 0, nil
}
