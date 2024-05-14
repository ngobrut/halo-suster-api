package custom_error

import (
	"net/http"

	"github.com/ngobrut/halo-sus-api/constant"
)

type CustomError struct {
	ErrorContext *ErrorContext
}

type ErrorContext struct {
	HTTPCode int
	Message  string
}

func (c *CustomError) Error() string {
	if c.ErrorContext.HTTPCode == 0 {
		c.ErrorContext.HTTPCode = http.StatusInternalServerError
	}

	return constant.HTTPStatusText(http.StatusInternalServerError)
}

func SetCustomError(errContext *ErrorContext) *CustomError {
	return &CustomError{
		ErrorContext: errContext,
	}
}
