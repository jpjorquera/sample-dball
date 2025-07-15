package repository

import "dballz/internal/dto"

type CharacterExternalAPIRepository struct {
	apiURL string
}

func NewCharacterExternalRepository(apiURL string) *CharacterExternalAPIRepository {
	return &CharacterExternalAPIRepository{apiURL: apiURL}
}

func (r *CharacterExternalAPIRepository) GetByName(name string) (*dto.CharacterInformation, error) {
	return nil, nil
}
