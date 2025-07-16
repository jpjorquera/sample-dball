package repository

import (
	"dballz/internal/model"
)

type CharacterGetter interface {
	GetByName(name string) (*model.Character, error)
}

type CharacterSaver interface {
	Save(character *model.Character) error
}

type CharacterStore interface {
	CharacterGetter
	CharacterSaver
}
