package repository

import (
	"dballz/internal/model"
	"dballz/internal/repository/entities"
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
	var characterEntity entities.CharacterEntity
	err := r.db.Where("name LIKE ?", "%"+name+"%").First(&characterEntity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	characterDTO := &model.Character{
		ID:          characterEntity.ID,
		ExternalID:  characterEntity.ExternalID,
		Name:        characterEntity.Name,
		Race:        characterEntity.Race,
		Ki:          characterEntity.Ki,
		Description: characterEntity.Description,
	}
	return characterDTO, nil
}

func (r *CharacterDBRepository) Save(characterDTO *model.Character) error {
	characterEntity := entities.CharacterEntity{
		ExternalID:  characterDTO.ExternalID,
		Name:        characterDTO.Name,
		Race:        characterDTO.Race,
		Ki:          characterDTO.Ki,
		Description: characterDTO.Description,
	}

	err := r.db.Save(&characterEntity).Error
	if err != nil {
		return err
	}

	characterDTO.ID = characterEntity.ID
	return nil
}
