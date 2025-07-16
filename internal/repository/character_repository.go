package repository

import "dballz/internal/dto"

type CharacterGetter interface {
	GetByName(name string) (*dto.CharacterInformation, error)
}

type CharacterSaver interface {
	Save(character *dto.CharacterInformation) error
}

type CharacterStore interface {
	CharacterGetter
	CharacterSaver
}
