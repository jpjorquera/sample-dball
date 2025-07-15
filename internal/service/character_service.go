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

func (s *CharacterService) GenerateCharacter(name string) (dto.CharacterInformation, error) {
	characterInformation, err := s.externalGetter.GetByName(name)
	if err != nil {
		return dto.CharacterInformation{}, err
	}
	return characterInformation, nil
}
