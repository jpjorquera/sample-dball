package repository

import (
	"dballz/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CharacterExternalResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Ki          string `json:"ki"`
	Description string `json:"description"`
}

type CharacterExternalAPIRepository struct {
	apiURL string
}

func NewCharacterExternalRepository(apiURL string) *CharacterExternalAPIRepository {
	return &CharacterExternalAPIRepository{apiURL: apiURL}
}

func (r *CharacterExternalAPIRepository) GetByName(name string) (*model.Character, error) {
	url := fmt.Sprintf("%s/characters?name=%s", r.apiURL, name)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var characters []CharacterExternalResponse
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(characters) == 0 {
		return nil, ErrNotFound
	}

	foundCharacter := characters[0]
	return &model.Character{
		ExternalID:  foundCharacter.ID,
		Name:        foundCharacter.Name,
		Race:        foundCharacter.Race,
		Ki:          foundCharacter.Ki,
		Description: foundCharacter.Description,
	}, nil
}
