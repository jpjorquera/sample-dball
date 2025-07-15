package dto

type CreateCharacterRequest struct {
	Name string `json:"name" binding:"required"`
}

type CharacterInformation struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Ki          string `json:"ki"`
	Description string `json:"description"`
}
