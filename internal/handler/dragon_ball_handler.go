package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterDragonBallHandler(r *gin.Engine) {
	r.POST("/dragon-ball/character", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Id": "placeholder response",
		})
	})
}
