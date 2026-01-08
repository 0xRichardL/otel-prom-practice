package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppErrorType string

const (
	AppErrorTypeValidation AppErrorType = "validation"
	AppErrorTypeInternal   AppErrorType = "internal"
)

type ApplicationError struct {
	Message string
	Type    AppErrorType
}

func NewApplicationError(message string, errType AppErrorType) *ApplicationError {
	return &ApplicationError{
		Message: message,
		Type:    errType,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	statusMap = map[AppErrorType]int{
		AppErrorTypeValidation: http.StatusBadRequest,
		AppErrorTypeInternal:   http.StatusInternalServerError,
	}
)

func RespondError(c *gin.Context, err error) {
	var appErr *ApplicationError
	if ok := errors.As(err, &appErr); ok {
		code := statusMap[appErr.Type]
		c.JSON(code, ErrorResponse{
			Error:   http.StatusText(code),
			Code:    code,
			Message: appErr.Message,
		})
		return
	}

	// Default to internal server error if not an ApplicationError
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Error:   http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
		Message: err.Error(),
	})
}
