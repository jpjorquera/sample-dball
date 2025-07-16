package service

import (
	"errors"
	"testing"

	"dballz/internal/dto"
	"dballz/internal/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockStore struct{ mock.Mock }

func (m *mockStore) GetByName(name string) (*dto.CharacterInformation, error) {
	args := m.Called(name)
	if info, ok := args.Get(0).(*dto.CharacterInformation); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *mockStore) Save(info *dto.CharacterInformation) error {
	args := m.Called(info)
	return args.Error(0)
}

type mockExternalGetter struct{ mock.Mock }

func (m *mockExternalGetter) GetByName(name string) (*dto.CharacterInformation, error) {
	args := m.Called(name)
	if info, ok := args.Get(0).(*dto.CharacterInformation); ok {
		return info, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestReturnCharacterFromStore(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &dto.CharacterInformation{ID: 1, ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	service := NewCharacterService(externalApi, store)

	gottenCharacter, err := service.GenerateCharacter(requestedCharacterName)

	require.NoError(t, err)
	assert.Equal(t, expectedCharacter, gottenCharacter)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestGetExternalCharacterAndSaves(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &dto.CharacterInformation{ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Gohan"

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	store.On("Save", expectedCharacter).Return(nil)
	service := NewCharacterService(externalApi, store)

	gottenCharacter, err := service.GenerateCharacter(requestedCharacterName)

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
	service := NewCharacterService(externalApi, store)

	gottenCharacter, err := service.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, ErrDatabase, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestReturnsErrorWhenExternalAPIFails(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	requestedCharacterName := "Goku"
	mockExternalError := errors.New("mock external error")

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(nil, mockExternalError)

	service := NewCharacterService(externalApi, store)

	gottenCharacter, err := service.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, mockExternalError, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}

func TestReturnsDatabaseErrorWhenStoreSaveFails(t *testing.T) {
	store := new(mockStore)
	externalApi := new(mockExternalGetter)
	expectedCharacter := &dto.CharacterInformation{ExternalID: 1, Name: "Goku", Race: "Saiyan", Description: "Sample description"}
	requestedCharacterName := "Goku"

	store.On("GetByName", requestedCharacterName).Return(nil, repository.ErrNotFound)
	externalApi.On("GetByName", requestedCharacterName).Return(expectedCharacter, nil)
	store.On("Save", expectedCharacter).Return(errors.New("mock database failure"))
	service := NewCharacterService(externalApi, store)

	gottenCharacter, err := service.GenerateCharacter(requestedCharacterName)

	assert.Nil(t, gottenCharacter)
	assert.Equal(t, ErrDatabase, err)
	store.AssertExpectations(t)
	externalApi.AssertExpectations(t)
}
