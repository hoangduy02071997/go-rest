package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewFullErrorResponse(statusCode int, rootError error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    rootError,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewErrorResponse(rootError error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    rootError,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnAuthorizedResponse(rootError error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    rootError,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// Sử dụng Decorator Pattern
// Mục đích ẩn lỗi ứng dụng (lỗi sql ...)
// Error ( Error ( Error () ) )
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

// Để customer error chỉ cần implement interface Error() string
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something when wrong with DB", err.Error(), "DB_ERROR")
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot not create %s", strings.ToLower(entity)),
		fmt.Sprintf("CannotCreateEntity%s", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Not found %s", strings.ToLower(entity)),
		fmt.Sprintf("EntityNotFound%s", entity),
	)
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid Request", err.Error(), "ERRINVALIDREQUEST")
}

func ErrInternalServer(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something when wrong in server",
		err.Error(),
		"Internal",
	)
}

func ErrParseRequest(err error) *AppError {
	return NewErrorResponse(err, "Failed to parse request", err.Error(), "PARE_REQUEST")
}
