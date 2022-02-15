package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrEmailAlreadyExists = "User with given email already exists"
)

var (
	errInternalServer = errors.New("internal server error")
)

// Rest error interface
type IRestError interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

// Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

// New Rest Error
func NewRestError(status int, err string, causes interface{}) IRestError {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

// NewInternalServerError:
func NewInternalServerError(causes interface{}) IRestError {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  errInternalServer.Error(),
		ErrCauses: causes,
	}
}

// ParseErrors
func ParseErrors(err error) IRestError {
	switch {
	case strings.Contains(strings.ToLower(err.Error()), "failed creating"):
		return NewRestError(http.StatusBadRequest, "Failed creating entity record", err)

	case strings.Contains(strings.ToLower(err.Error()), "failed updating"):
		return NewRestError(http.StatusBadRequest, "Failed updating entity record", err)

	case strings.Contains(strings.ToLower(err.Error()), "failed deleting"):
		return NewRestError(http.StatusBadRequest, "Failed deleting entity record", err)

	case strings.Contains(strings.ToLower(err.Error()), "failed querying"):
		return NewRestError(http.StatusBadRequest, "Failed querying entity record", err)

	default:
		if restErr, ok := err.(RestError); ok {
			return restErr
		}

		return NewInternalServerError(err)
	}
}

// ErrorResponse:
func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
