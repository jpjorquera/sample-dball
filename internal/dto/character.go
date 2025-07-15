package dto

type CreateCharacterRequest struct {
	Name string `json:"name" binding:"required"`
}
