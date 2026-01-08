package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (s *App) respondError(c *gin.Context, message string, code int) {
	err := ErrorResponse{
		Error:   http.StatusText(code),
		Code:    code,
		Message: message,
	}

	c.JSON(code, err)
}
