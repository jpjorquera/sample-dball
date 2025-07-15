package handler

import (
	"dballz/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userService *service.UserService) {
	r.GET("/users", func(c *gin.Context) {
		users, err := userService.ListUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})
}
