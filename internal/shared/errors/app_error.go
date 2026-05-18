package errors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Code       string
	Message    string
	StatusCode int
	Err        error
}

const (
	ErrInternal     = "INTERNAL_ERROR"
	ErrNotFound     = "NOT_FOUND"
	ErrInvalidInput = "INVALID_INPUT"
	ErrConflict     = "CONFLICT"
	ErrUnauthorized = "UNAUTHORIZED"
	ErrForbidden    = "FORBIDDEN"
)

func (e *AppError) Error() string {
	return e.Message
}

func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

func New(code, message string, status int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: status,
		Err:        err,
	}
}

func Internal(err error) *AppError {
	return New(
		ErrInternal,
		"internal server error",
		http.StatusInternalServerError,
		err,
	)
}

func NotFound(message string, err error) *AppError {
	return New(
		ErrNotFound,
		message,
		http.StatusNotFound,
		err,
	)
}

func InvalidInput(message string, err error) *AppError {
	return New(
		ErrInvalidInput,
		message,
		http.StatusBadRequest,
		err,
	)
}

func Conflict(message string, err error) *AppError {
	return New(
		ErrConflict,
		message,
		http.StatusConflict,
		err,
	)
}

func Unauthorized(message string, err error) *AppError {
	return New(
		ErrUnauthorized,
		message,
		http.StatusUnauthorized,
		nil,
	)
}

func Forbidden(message string) *AppError {
	return New(
		ErrForbidden,
		message,
		http.StatusForbidden,
		nil,
	)
}