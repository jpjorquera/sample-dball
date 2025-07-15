package service

import (
	"dballz/internal/dto"
	"dballz/internal/repository"
)

type CharacterService struct {
	externalGetter repository.CharacterGetter
	store          repository.CharacterStore
}

func NewCharacterService(extGetter repository.CharacterGetter, store repository.CharacterStore) *CharacterService {
	return &CharacterService{
		externalGetter: extGetter,
		store:          store,
	}
}

func (s *CharacterService) GenerateCharacter(dto.CreateCharacterRequest) (int, error) {
	return 0, nil
}
