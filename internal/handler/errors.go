package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Message string `json:"error"`
	Code    int    `json:"-"`
}

func (e APIError) Respond(c *gin.Context) {
	c.JSON(e.Code, e)
}

var (
	ErrBadRequest = APIError{Message: "Invalid request", Code: http.StatusBadRequest}
	ErrNotFound   = APIError{Message: "Resource not found", Code: http.StatusNotFound}
	ErrServer     = APIError{Message: "Internal server error", Code: http.StatusInternalServerError}
	ErrUpstream   = APIError{Message: "upstream service unavailable", Code: http.StatusBadGateway}
)
