package repository_test

import (
	"testing"

	"dballz/internal/model"
	"dballz/internal/repository"
	"dballz/internal/repository/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(entities.AllModels()...))
	return db
}

func TestSaveAndGetCharacter(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewCharacterDBRepository(db)

	character := &model.Character{
		ExternalID: 1,
		Name:       "Goku",
		Race:       "Saiyan",
		Ki:         "9000",
	}
	err := repo.Save(character)
	require.NoError(t, err)
	assert.NotZero(t, character.ID)

	gottenCharacter, err := repo.GetByName("Goku")
	require.NoError(t, err)
	assert.Equal(t, "Goku", gottenCharacter.Name)
	assert.Equal(t, "Saiyan", gottenCharacter.Race)
}

func TestSaveAndGetCharacterByName(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewCharacterDBRepository(db)

	character := &model.Character{
		ExternalID: 1,
		Name:       "Goku",
		Race:       "Saiyan",
		Ki:         "9000",
	}
	err := repo.Save(character)
	require.NoError(t, err)
	assert.NotZero(t, character.ID)

	gottenCharacter, err := repo.GetByName("Goku")
	require.NoError(t, err)
	assert.Equal(t, "Goku", gottenCharacter.Name)
	assert.Equal(t, "Saiyan", gottenCharacter.Race)
}

func TestGetByNameReturnsNotFoundWhenMissing(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewCharacterDBRepository(db)

	gottenCharacter, err := repo.GetByName("Goku")
	assert.Error(t, err)
	assert.Equal(t, repository.ErrNotFound, err)
	assert.Nil(t, gottenCharacter)
}
