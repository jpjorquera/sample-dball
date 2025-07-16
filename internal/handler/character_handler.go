package handler

import (
	"errors"
	"net/http"

	"dballz/internal/model"
	"dballz/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterCharacterHandler(r *gin.Engine, characterService *service.CharacterService) {
	r.POST("/dragon-ball/character", func(c *gin.Context) {
		var req model.CreateCharacterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrBadRequest)
			return
		}

		character, err := characterService.GenerateCharacter(req.Name)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrNotFound):
				ErrNotFound.Respond(c)
			case errors.Is(err, service.ErrExternalAPI):
				ErrUpstream.Respond(c)
			default:
				ErrServer.Respond(c)
			}
			return
		}
		c.JSON(http.StatusOK, character)
	})
}
