package model

type CreateCharacterRequest struct {
	Name string `json:"name" binding:"required"`
}

type Character struct {
	ID          uint   `json:"id"`
	ExternalID  uint   `json:"external_id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Ki          string `json:"ki"`
	Description string `json:"description"`
}
