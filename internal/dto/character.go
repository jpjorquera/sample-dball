package dto

type CreateCharacterRequest struct {
	Name string `json:"name" binding:"required"`
}

type CharacterInformation struct {
	ID          int
	Name        string
	Ki          string
	Race        string
	Description string
}
