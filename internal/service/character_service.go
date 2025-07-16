package service

import (
	"dballz/internal/dto"
	"dballz/internal/repository"
	"errors"
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

func (s *CharacterService) GenerateCharacter(name string) (*dto.CharacterInformation, error) {
	characterInformation, err := s.store.GetByName(name)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, ErrDatabase
	}

	if characterInformation != nil {
		return characterInformation, nil
	}

	characterInformation, err = s.externalGetter.GetByName(name)
	if err != nil {
		return nil, err
	}

	err = s.store.Save(characterInformation)
	if err != nil {
		return nil, ErrDatabase
	}

	return characterInformation, nil
}
