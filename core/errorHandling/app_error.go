package errorHandling

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	Code    int    `json:"code"`
	RootErr error  `json:"-"`
	Message string `json:"message"`
	Log     string `json:"log"`
	Key     string `json:"key"`
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		Code:    http.StatusBadRequest,
		RootErr: root,
		Message: msg,
		Log:     log,
		Key:     key,
	}
}

func NewFullErrorResponse(code int, root error, msg, log, key string) *AppError {
	return &AppError{
		Code:    code,
		RootErr: root,
		Message: msg,
		Log:     log,
		Key:     key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		Code:    http.StatusUnauthorized,
		RootErr: root,
		Message: msg,
		Key:     key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func ErrUnauthorized(err error) *AppError {
	if err != nil {
		return NewUnauthorized(err, err.Error(), "INVALID_TOKEN")
	}
	return NewUnauthorized(err, "token không hợp lệ", "INVALID_TOKEN")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}
