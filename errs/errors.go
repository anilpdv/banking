package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message}
}

// NewNotFoundError : Not Found Error Helper Method
func NewNotFoundError(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

// NewUnexpectedError : Not Found Error Helper Method
func NewUnexpectedError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}
