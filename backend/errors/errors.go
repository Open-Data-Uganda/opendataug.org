package errors

import "net/http"

type ErrorType string

const (
	ErrorTypeValidation   ErrorType = "VALIDATION_ERROR"
	ErrorTypeDatabase     ErrorType = "DATABASE_ERROR"
	ErrorTypeNotFound     ErrorType = "NOT_FOUND"
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"
	ErrorTypeBadRequest   ErrorType = "BAD_REQUEST"
	ErrorTypeInternal     ErrorType = "INTERNAL_ERROR"
	ErrorTypeRateLimit    ErrorType = "RATE_LIMIT_ERROR"
)

type APIError struct {
	Type       ErrorType `json:"type"`
	Message    string    `json:"message"`
	Details    any       `json:"details,omitempty"`
	StatusCode int       `json:"status_code"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewValidationError(message string, details any) *APIError {
	return &APIError{
		Type:       ErrorTypeValidation,
		Message:    message,
		Details:    details,
		StatusCode: http.StatusBadRequest,
	}
}

func NewDatabaseError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeDatabase,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeNotFound,
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}

func NewBadRequestError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeBadRequest,
		Message:    message,
		StatusCode: http.StatusBadRequest,
	}
}

func NewInternalError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeInternal,
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}

func NewRateLimitError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeRateLimit,
		Message:    message,
		StatusCode: http.StatusTooManyRequests,
	}
}

func NewUnauthorizedError(message string) *APIError {
	return &APIError{
		Type:       ErrorTypeUnauthorized,
		Message:    message,
		StatusCode: http.StatusUnauthorized}
}
