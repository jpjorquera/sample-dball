package repository

import (
	"dballz/internal/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CharacterExternalAPIRepository struct {
	apiURL string
}

func NewCharacterExternalRepository(apiURL string) *CharacterExternalAPIRepository {
	return &CharacterExternalAPIRepository{apiURL: apiURL}
}

func (r *CharacterExternalAPIRepository) GetByName(name string) (*dto.CharacterInformation, error) {
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

	var characters []dto.CharacterInformation
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(characters) == 0 {
		return nil, ErrNotFound
	}
	return &characters[0], nil
}
