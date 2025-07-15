package service

import (
	"dballz/internal/dto"
	"dballz/internal/repository"
)

type CharacterService struct {
	repo *repository.CharacterRepository
}

func NewCharacterService(repo *repository.CharacterRepository) *CharacterService {
	return &CharacterService{repo: repo}
}

func (s *CharacterService) GenerateCharacter(dto.CreateCharacterRequest) (int, error) {
	return 0, nil
}
