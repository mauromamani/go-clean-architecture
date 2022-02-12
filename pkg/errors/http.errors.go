package errors

import (
	"errors"
	"fmt"
	"net/http"
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

// NewInternalServerError:
func NewInternalServerError(causes interface{}) IRestError {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  errInternalServer.Error(),
		ErrCauses: causes,
	}
}
