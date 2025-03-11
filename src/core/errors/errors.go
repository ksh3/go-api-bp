package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Cause      error  `json:"-"`
	IsCritical bool   `json:"-"`
}

func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// NOTE: Soft errros, code range: 7000-19999
func ValidationError(appErrorCode int, message string, cauese error) *AppError {
	return &AppError{
		Code:       appErrorCode,
		Message:    message,
		Cause:      cauese,
		IsCritical: false,
	}
}

// NOTE: Hard errors, code range: http status code range(200-599)
func InternalError(message string, cauese error) *AppError {
	return &AppError{
		Code:       http.StatusInternalServerError,
		Message:    message,
		Cause:      cauese,
		IsCritical: true,
	}
}
