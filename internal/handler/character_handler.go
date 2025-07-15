package handler

import (
	"net/http"

	"dballz/internal/dto"
	"dballz/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterCharacterHandler(r *gin.Engine, characterService *service.CharacterService) {
	r.POST("/dragon-ball/character", func(c *gin.Context) {
		var req dto.CreateCharacterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"details": "Invalid value for field name",
			})
			return
		}

		character, err := characterService.GenerateCharacter(req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, character)
	})
}
