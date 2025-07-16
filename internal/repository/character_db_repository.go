package repository

import (
	"dballz/internal/model"
	"errors"

	"gorm.io/gorm"
)

type CharacterDBRepository struct {
	db *gorm.DB
}

func NewCharacterDBRepository(database *gorm.DB) *CharacterDBRepository {
	return &CharacterDBRepository{db: database}
}

func (r *CharacterDBRepository) GetByName(name string) (*model.Character, error) {
	var characterModel model.Character
	err := r.db.Where("name LIKE ?", "%"+name+"%").First(&characterModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	characterDTO := &model.Character{
		ID:          characterModel.ID,
		ExternalID:  characterModel.ExternalID,
		Name:        characterModel.Name,
		Race:        characterModel.Race,
		Ki:          characterModel.Ki,
		Description: characterModel.Description,
	}
	return characterDTO, nil
}

func (r *CharacterDBRepository) Save(character *model.Character) error {
	characterModel := model.Character{
		ExternalID:  character.ExternalID,
		Name:        character.Name,
		Race:        character.Race,
		Ki:          character.Ki,
		Description: character.Description,
	}

	err := r.db.Save(&characterModel).Error
	if err != nil {
		return err
	}

	character.ID = characterModel.ID
	return nil
}
