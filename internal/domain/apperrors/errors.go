package apperrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Dann-Go/InnoTaxiDriverService/internal/domain/responses"

	"github.com/gin-gonic/gin"
)

var (
	ErrPhoneIsAlreadyTaken  = errors.New("driver with such phone already exists")
	ErrEmailIsAlreadyTaken  = errors.New("driver with such email already exists")
	ErrInternalServer       = errors.New("internal server error")
	ErrNoRecords            = errors.New("no records")
	ErrBadRequest           = errors.New("incorrect syntax input")
	ErrUserNotFound         = errors.New("user not found")
	ErrTokenInvalid         = errors.New("invalid token")
	ErrWrongPassword        = errors.New("wrong password")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrWrongTokenClaims     = errors.New("token claims are not of type *tokenClaims")
)

type HTTPError struct {
	StatusCode int
	Message    string
}

func ErrorResponse(context *gin.Context, err error) {

	newErr := UnWrapper(err)
	response := newErrorResponse(newErr)
	context.AbortWithStatusJSON(response.StatusCode, responses.ServerResponse{
		Success: false,
		Msg:     response.Message,
	})
}

func Wrapper(message error, err error) error {
	wrap := fmt.Errorf("%w: %s", message, err)
	return wrap
}

func UnWrapper(err error) error {
	e := errors.Unwrap(err)
	if e != nil {
		return e
	}
	return err
}

func newErrorResponse(err error) HTTPError {
	var httpErr HTTPError
	switch {
	case errors.Is(err, ErrPhoneIsAlreadyTaken):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrEmailIsAlreadyTaken):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrTokenInvalid):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrUserNotFound):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrBadRequest):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrInternalServer):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrNoRecords):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrWrongPassword):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrInvalidSigningMethod):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	case errors.Is(err, ErrWrongTokenClaims):
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusBadRequest
	default:
		httpErr.Message = err.Error()
		httpErr.StatusCode = http.StatusInternalServerError
	}

	return httpErr
}
