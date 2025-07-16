package service

import (
	"errors"
	"testing"

	"dballz/internal/model"
	"dballz/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockStore struct{ mock.Mock }

func (m *mockStore) GetByName(name string) (*model.Character, error) {
	args := m.Called(name)
	if info, ok := args.Get(0).(*model.Character); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockStore) Save(info *model.Character) error {
	args := m.Called(info)
	return args.Error(0)
}

type mockExternalGetter struct{ mock.Mock }

func (m *mockExternalGetter) GetByName(name string) (*model.Character, error) {
	args := m.Called(name)
	if info, ok := args.Get(0).(*model.Character); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestReturnCharacterFromStore(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &model.Character{ID: 1, ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	characterService := NewCharacterService(externalApi, store)

	gottenCharacter, err := characterService.GenerateCharacter(requestedCharacterName)

	require.NoError(t, err)
	assert.Equal(t, expectedCharacter, gottenCharacter)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestGetExternalCharacterAndSaves(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &model.Character{ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Gohan"

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	store.On("Save", expectedCharacter).Return(nil)
	characterService := NewCharacterService(externalApi, store)

	gottenCharacter, err := characterService.GenerateCharacter(requestedCharacterName)

	require.NoError(t, err)
	assert.Equal(t, expectedCharacter, gottenCharacter)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestReturnsDatabaseErrorWhenStoreGetFails(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(nil, errors.New("mock error"))
	characterService := NewCharacterService(externalApi, store)

	gottenCharacter, err := characterService.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, ErrDatabase, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestReturnsErrorWhenExternalAPIFails(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(nil, ErrExternalAPI)
	characterService := NewCharacterService(externalApi, store)

	gottenCharacter, err := characterService.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, ErrExternalAPI, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestReturnsDatabaseErrorWhenStoreSaveFails(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &model.Character{ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	store.On("Save", expectedCharacter).Return(errors.New("mock database failure"))
	characterService := NewCharacterService(externalApi, store)

	gottenCharacter, err := characterService.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, ErrDatabase, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}
