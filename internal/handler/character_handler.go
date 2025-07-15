package handler

import (
	"net/http"

	"dballz/internal/dto"

	"github.com/gin-gonic/gin"
)

func RegisterCharacterHandler(r *gin.Engine) {
	r.POST("/dragon-ball/character", func(c *gin.Context) {
		var req dto.CreateCharacterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
				"details": "Invalid value for field name",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Id": "placeholder response",
		})
	})
}
