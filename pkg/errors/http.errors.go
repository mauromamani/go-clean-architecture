package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/mauromamani/go-clean-architecture/pkg/validator"
)

var (
	errInternalServer     = errors.New("internal server error")
	errRecordNotFound     = errors.New("record not found")
	errBadRequest         = errors.New("bad request")
	errDuplicateEmail     = errors.New("duplicate email")
	errInvalidIDParameter = errors.New("invalid ID parameter")
	errInvalidForeignKey  = errors.New("invalid foreign key")
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
	ErrError  interface{} `json:"error,omitempty"`
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
func NewRestError(status int, err interface{}, causes interface{}) IRestError {
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

// NewBadRequestError:
func NewBadRequestError(causes interface{}) IRestError {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  errBadRequest.Error(),
		ErrCauses: causes,
	}
}

// ParseErrors
func ParseErrors(err error) IRestError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, errRecordNotFound.Error(), err)

	case strings.Contains(strings.ToLower(err.Error()), "invalid id parameter"):
		return NewRestError(http.StatusBadRequest, errInvalidIDParameter.Error(), err)

	case strings.Contains(strings.ToLower(err.Error()), "duplicate key value violates unique constraint \"users_email_key\""):
		return NewRestError(http.StatusUnprocessableEntity, errDuplicateEmail.Error(), err)

	case strings.Contains(strings.ToLower(err.Error()), "insert or update on table \"posts\" violates foreign key constraint \"posts_user_id_fkey\""):
		return NewRestError(http.StatusUnprocessableEntity, errInvalidForeignKey.Error(), err)

	case strings.Contains(strings.ToLower(err.Error()), "field validation"):
		return parseValidatorError(err)

	default:
		if restErr, ok := err.(RestError); ok {
			return restErr
		}

		return NewInternalServerError(err)
	}
}

func parseValidatorError(err error) IRestError {
	errors := validator.MapTranslatedErrors(err)
	return NewRestError(http.StatusUnprocessableEntity, errors, err)
}

// ErrorResponse:
func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
