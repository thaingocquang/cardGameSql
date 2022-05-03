package common

import (
	"errors"
	"net/http"
)

// AppError ...
type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"-"`       // not send to client
	Message    string `json:"message"` // readable for user
	Log        string `json:"log"`     // readable for system
	Key        string `json:"key"`     // error key
}

// NewErrorResponse ...
func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// NewFullErrorResponse ...
func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// NewUnAuthorized ...
func NewUnAuthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

// NewCustomError ...
func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return &AppError{
			RootErr: root,
			Message: msg,
			Log:     root.Error(),
			Key:     key,
		}
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

// RootError ...
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootErr()
	}
	return e.RootErr
}

// Error implements Error interface
func (e *AppError) Error() string {
	return e.RootError().Error()
}

// ErrDB ...
func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

// ErrInternal ...
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "ErrInternal")
}
